/* SMOOTHLIFE (STEPHAN RAFLER)
criado em       : 2024/11/20
ult. atualização: 2024/12/28
autor           : Beatriz Bittencourt <beatrizdecbittencourt@gmail.com>
notas           : Executa o SmoothLife (parâmetros originais) em arquivos .dat
compilação      : -
execução        : go run smoothlife.go
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

// constantes como dimensões x e y da rede (Nx e Ny), número de gerações (NG), parâmetros para geração da condição inicial e parâmetros do SL
const (
	Nx, Ny, NG, n_q, max_q, ra, alphaN, alphaM, b1, b2, d1, d2, dt, ri = 500, 500, 500, 30, 50, 21.0, 0.028, 0.147, 0.278, 0.365, 0.267, 0.445, 1.0, ra / 3
)

// variáveis como x e y da célula, número do arquivo .dat, número da célula e contagens de células, gerações e arquivos
var (
	i, j, x_q, y_q, num, t, n int
	phi [Nx * Ny]float64 // float64 para assumir estados decimais
	updt [Nx * Ny]float64 // idem
	dat [2]int
)

// função condição inicial, distribui aleatoriamente quadrados com células entre 0 e 1
func ic() {
	for n = 0; n < n_q; n++ { // cria quadrados até atingir uma quantidade máxima (n_q)
		x_q = rand.Intn(Ny - max_q) // coordenadas iniciais de cada quadrado (x e y)
		y_q = rand.Intn(Nx - max_q)
		for i = y_q; i < y_q+max_q; i++ { // na área dentro do perímetro...
			for j = x_q; j < x_q+max_q; j++ {
				phi[i*Ny+j] = rand.Float64() // ...células assumem estados aleatórios entre 0 e 1
			}
		}
	}
}

// função op, imprime a rede ao final de cada geração
func op(num int, phi [Nx * Ny]float64) {
	f := fmt.Sprintf("smoothlife-%v.dat", num) // arquivo smoothlife-tcc-(x).dat (a depender do número da geração)
	file, _ := os.Create(f) // cria o arquivo
	for i = 0; i < Nx; i++ { // para toda a rede
		for j = 0; j < Ny; j++ {
			fmt.Fprintf(file, "%f%f ", phi[j*Nx+i], phi[j*Nx+i]) // imprime o valor de cada célula em ordem, duas vezes
		}
		fmt.Fprintf(file, "\n") // pula linha ao atingir Ny
	}
	defer file.Close()
}

// equações do SmoothLife abaixo; todas são encontradas no artigo original (Stephan Rafler)

// função sigma
func sigma(x, a, alpha float64) float64 { // note abaixo que alpha toma os valores de alphaN e alphaM
	return 1.0 / (1.0 + float64(math.Exp(-float64((x-a)*4/alpha))))
}

// função sigmaN (depende de sigma)
func sigmaN(x, a, b float64) float64 {
	return sigma(x, a, alphaN) * (1 - sigma(x, b, alphaN))
}

// função sigmaM (depende de sigma)
func sigmaM(x, y, m float64) float64 {
	return x*(1-sigma(m, 0.5, alphaM)) + y*sigma(m, 0.5, alphaM)
}

// função de transição s(n, m) (depende de sigmaN e sigmaM)
func s(n, m float64) float64 {
	return sigmaN(n, sigmaM(b1, d1, m), sigmaM(b2, d2, m))
}

// função principal - a simulação ocorre essencialmente aqui
func main() {
	rand.Seed(time.Now().UnixNano())
	ic() // ativa a condição inicial
	op(0, phi) // ativa op para imprimir a condição inicial

	for i := 0; i < Nx*Ny; i++ { // calcula a densidade populacional de cada estado
		if phi[i] > 0.5 { // condição: vivo se > 0,5 e morto se < 0,5
			dat[1]++
		} else {
			dat[0]++
		}
	}
	f, _ := os.Create("pop-smoothlife.dat") // cria o arquivo .dat de densidade populacional
	defer f.Close()
	fmt.Fprintf(f, "%.2f   %.2f\n", float64(dat[1])/(Nx*Ny), float64(dat[0])/(Nx*Ny)) // imprime as densidades em %
	dat[1] = 0 // zera as densidades para serem atualizadas a cada geração
	dat[0] = 0

	for t = 1; t <= NG; t++ { // para cada geração
		for i = 0; i < Nx; i++ { // para toda a rede
			for j = 0; j < Ny; j++ {
				m, M := float64(0), float64(0) // m, M, n e N começam em 0,0
				n, N := float64(0), float64(0)

				for dy := -(ra - 1); dy <= (ra - 1); dy++ { // calcula preenchimentos da vizinhança com condições de contorno
					for dx := -(ra - 1); dx <= (ra - 1); dx++ {
						x := (int(float64(i) + dx + Nx)) % Nx
						y := (int(float64(j) + dy + Ny)) % Ny
						if dx*dx+dy*dy <= ri*ri { // vizinhança interna, preenchimento m
							m += phi[y*Nx+x]
							M++
						} else if dx*dx+dy*dy <= ra*ra { // vizinhança externa, preenchimento n
							n += phi[y*Nx+x]
							N++
						}
					}
				}
				m /= M
				n /= N
				q := s(n, m) // agora que tem n e m, calcula a função de transição
				updt[j*Nx+i] = 2*q - 1 // atualização da rede
			}
		}
		for i = 0; i < Nx; i++ { // para toda a rede
			for j = 0; j < Ny; j++ {
				phi[j*Nx+i] += math.Min(math.Max(dt * updt[j*Nx+i], 0), 1) // troca a rede inicial para a atualizada
				if phi[j*Nx+i] > 0.5 { // atualiza a densidade populacional
					dat[1]++
				} else {
					dat[0]++
				}
			}
		}
		fmt.Fprintf(f, "%.2f   %.2f\n", float64(dat[1])/(Nx*Ny), float64(dat[0])/(Nx*Ny)) // imprime a densidade populacional atualizada em %                                                                              // aumenta a cada geração
		op(t, phi) // cria o arquivo .dat a cada geração (smoothlife-tcc-1, smoothlife-tcc-2...)
		fmt.Printf("%d/%d\n", t, NG) // mostra gerações completas em tempo real no terminal
		dat[1] = 0 // zera as densidades populacionais novamente
		dat[0] = 0
	}
}
