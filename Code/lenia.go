/* LENIA (BERT WANG-CHAK CHAN)
criado em       : 2024/12/10
ult. atualização: 2026/01/06
autor           : Beatriz Bittencourt <beatrizdecbittencourt@gmail.com>
notas           : Executa Lenia (funções e parâmetros ***para Orbium unicaudatus, Orbidae***) em arquivos .dat
compilação      : -
execução        : go run lenia.go
*/

package main

// bibliotecas utilizadas
import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

/* constantes: dimensões x e y da rede (Nx e Ny), número de gerações (NG), parâmetros para geração da condição inicial e parâmetros do Lenia

 n_q = número de quadrados da condição inicial | max_q = lado dos quadrados | R = raio de influência do kernel | alpha, mu, sigma = parâmetros | T = passos temporais por geração */
const (
	Nx, Ny, NG, n_q, max_q, R, alpha, mu, sigma, T = 200, 200, 200, 15, 18, 18.0, 4, 0.14, 0.015, 10.0
)

// variáveis: x e y da célula, x e y dos quadrados, número do arquivo .dat, tempo e contador geral de células, gerações e arquivos
var (
	i, j, x_q, y_q, n_arq, t, n int
	phi [Nx][Ny]float64 // estado da célula (2D por exigência do Lenia, e float64 para assumir estados decimais)
	dat [2]int
)

// função condição inicial, distribui aleatoriamente quadrados com células entre 0 e 1
func ic() {
	for n = 0; n < n_q; n++ { // cria quadrados até atingir uma quantidade máxima (n_q)
		x_q = rand.Intn(Nx - max_q) // coordenadas iniciais de cada quadrado (x e y)
		y_q = rand.Intn(Ny - max_q)
		for i = y_q; i < y_q+max_q; i++ { // na área dentro do perímetro...
			for j = x_q; j < x_q+max_q; j++ {
				phi[i][j] = math.Min(rand.Float64(), rand.Float64()) // ...células assumem estados aleatórios entre 0 e 1, favorecendo estados menores
			}
		}
	}
}

// função op, imprime a rede ao final de cada geração
func op(n_arq int, phi [Nx][Ny]float64) {
	f := fmt.Sprintf("lenia-%v.dat", n_arq) // arquivo lenia-(x).dat
	file, _ := os.Create(f) // cria o arquivo
	for i = 0; i < Nx; i++ { // para toda a rede
		for j = 0; j < Ny; j++ {
			fmt.Fprintf(file, "%f%f ", phi[i][j], phi[i][j]) // imprime o valor de cada célula em ordem, duas vezes
		}
		fmt.Fprintf(file, "\n") // pula linha ao atingir Ny
	}
	defer file.Close()
}

// função de convolução manual, dispensa bibliotecas extras
func convolucao(entrada1 [Nx][Ny]float64, entrada2 [][]float64, Nx, Ny int) [][]float64 { // entradas: matrizes que participam da convolução
	len_entrada2 := len(entrada2)
	offset := len_entrada2 / 2
	saida := make([][]float64, Nx) // matriz resultado da convolução
	for i := range saida {
		saida[i] = make([]float64, Ny)
	}

	for i := 0; i < Nx; i++ { // para toda célula
		for j := 0; j < Ny; j++ {
			var soma float64
			for ki := 0; ki < len_entrada2; ki++ { // condições de contorno
				for kj := 0; kj < len_entrada2; kj++ {
					x := (i + ki - offset + Nx) % Nx
					y := (j + kj - offset + Ny) % Ny
					soma += entrada1[x][y] * entrada2[ki][kj] // convolução em si
				}
			}
			saida[i][j] = soma
		}
	}

	return saida //função convolução retorna o resultado
}

// equações do Lenia abaixo; todas são encontradas no artigo original (Bert Wang-Chak Chan, 2019) e devem ser alteradas a depender do "espécime" de interesse (Orbidae)

// função Kernel, cria a vizinhança para toda célula
func Kernel(Nx, Ny int, R float64) [][]float64 {
	kernel := make([][]float64, Nx) // cria a matriz da vizinhança
	for i := range kernel {
		kernel[i] = make([]float64, Ny)
	}

	for i := 0; i < Nx; i++ { // para toda célula...
		for j := 0; j < Ny; j++ {
			dx := float64(i - Nx/2) // coordenadas da distância polar
			dy := float64(j - Ny/2)
			r_polar := math.Sqrt(dx*dx + dy*dy) / R // fórmula do raio polar
			if r_polar < 1 { // se o raio polar for menor que o raio da vizinhança...
				kernel[i][j] = math.Exp(alpha * (1 - (1.0/(4.0*r_polar*(1-r_polar))))) // ...cria a área de influência (nesse caso, apenas kernel core)
			}
		}
	}

	// normaliza a vizinhança 
	var soma float64
	for i := range kernel {
		for j := range kernel[i] {
			soma += kernel[i][j]
		}
	}
	for i := range kernel {
		for j := range kernel[i] {
			kernel[i][j] /= soma
		}
	}

	return kernel // função Kernel retorna a matriz da vizinhança
}

func g(U [][]float64, m, s float64) [][]float64 { // função G(u; mu, sigma) (growth mapping, mapeamento de crescimento)
	size := len(U) // tamanho da matriz G: tamanho da matriz de potencial U (que é a convolução)
	G := make([][]float64, size)
	for i := range G {
		G[i] = make([]float64, size)
		for j := range G[i] {
			G[i][j] = 2*(math.Exp(-(math.Pow((U[i][j]-mu), 2)) / (2 * sigma * sigma))) - 1 // fórmula da função G(u; mu, sigma) (Orbidae)
		}
	}
	return G // função g retorna a matriz mapeamento de crescimento G
}

// função principal - a simulação ocorre essencialmente aqui
func main() {
	rand.Seed(time.Now().UnixNano())
	ic() // ativa a condição inicial
	op(0, phi) // ativa op para imprimir a condição inicial

	K := Kernel(Nx, Ny, R) // inicia a vizinhança para todas as células

	for i := 0; i < Nx; i++ { // calcula a densidade populacional de cada estado
		for j := 0; j < Ny; j++ {
			if phi[i][j] > 0.5 { // condição: vivo se > 0,5 e morto se < 0,5
				dat[1]++
			} else {
				dat[0]++
			}
		}
	}

	f, _ := os.Create("pop-lenia.dat") // cria o arquivo .dat de densidade populacional
	defer f.Close()
	fmt.Fprintf(f, "%.2f   %.2f\n", float64(dat[1])/(Nx*Ny), float64(dat[0])/(Nx*Ny)) // imprime as densidades em %
	dat[1] = 0 // zera as densidades para serem atualizadas a cada geração
	dat[0] = 0

	for t := 1; t <= NG; t++ { // para cada geração
		U := convolucao(phi, K, Nx, Ny) // realiza a convolução do estado com a vizinhança
		G := g(U, mu, sigma) // calcula a função G(u; mu, sigma)

// note que kernel, convolucao e g já se aplicam para toda a rede; não é necessário usar laços for externamente

		for i := 0; i < Nx; i++ { // para toda a rede
			for j := 0; j < Ny; j++ {
				phi[i][j] = math.Min(math.Max(phi[i][j]+(1/T)*G[i][j], 0), 1)
				// troca a rede inicial para a atualizada, restringindo os valores
				if phi[i][j] > 0.5 { // atualiza a densidade populacional
					dat[1]++
				} else {
					dat[0]++
				}
			}
		}
		fmt.Fprintf(f, "%.2f   %.2f\n", float64(dat[1])/(Nx*Ny), float64(dat[0])/(Nx*Ny)) // imprime a densidade populacional atualizada em % 
		op(t, phi) // cria o arquivo .dat a cada geração (lenia-tcc-1, lenia-tcc-2...)
		fmt.Printf("%d/%d\n", t, NG) // mostra gerações completas em tempo real no terminal
		dat[1] = 0 // zera as densidades populacionais novamente
		dat[0] = 0
	}
}
