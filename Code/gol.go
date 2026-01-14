/* JOGO DA VIDA DE CONWAY
criado em       : 2024/10/31
ult. atualização: 2024/12/28
autor           : Beatriz Bittencourt <beatrizdecbittencourt@gmail.com>
notas           : Executa o Jogo da Vida de Conway (regra 23/3) em arquivos .dat
compilação      : -
execução        : go run gol.go
*/

package main

// bibliotecas utilizadas
import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// constantes como dimensões x e y da rede (Nx e Ny) e o número de gerações (NG)
const (
	Nx, Ny, NG = 150, 150, 1000
)

// variáveis como x e y da célula, número do arquivo .dat, número da célula e contagens de células, gerações e arquivos
var (
	i, j, num, ind, gd, t, vd, ve, vb, vc, vd1, vd2, vd3, vd4, soma int
	phi [Nx * Ny]int // rede
	updt [Nx * Ny]int // rede após atualizações
	dat [2]int // matriz de densidade populacional
)

// função condição inicial, distribui 0 ou 1 para cada célula sucessivamente
func ic() {
	for i = 0; i < Nx; i++ { // para todas as células da rede
		for j = 0; j < Ny; j++ {
			state := rand.Float64() // número pseudoaleatório entre 0 e 1
			if state < 0.5 { // probabilidade
				phi[i*Ny+j] = 1 // vivo
			} else {
				phi[i*Ny+j] = 0 // morto
			}
		}
	}
}

// função op, imprime a rede ao final de cada geração
func op(num int, phi [Nx * Ny]int) {
	f := fmt.Sprintf("gol-%v.dat", num) // arquivo jdv-tcc-0.dat (condição inicial)
	file, _ := os.Create(f) // cria o arquivo
	for i = 0; i < Nx; i++ { // para cada espaço no arquivo
		for j = 0; j < Ny; j++ {
			fmt.Fprintf(file, "%d ", phi[j*Nx+i]) // imprime o valor de cada célula em ordem
		}
		fmt.Fprintf(file, "\n") // pula linha ao atingir Ny
	}
	defer file.Close()
}

// função principal, Jogo da Vida ocorre aqui
func main() {
	rand.Seed(time.Now().UnixNano())
	ic() // ativa a condição inicial
	op(0, phi) // ativa op para imprimir a condição inicial

	for i = 0; i < Nx*Ny; i++ { // para toda a rede
		dat[phi[i]]++ // calcula a densidade populacional de cada estado
	}
	f, _ := os.Create("pop-gol.dat") // cria o arquivo .dat de densidade populacional
	fmt.Fprintf(f, "%.2f  %.2f\n", float64(dat[1])/(Nx*Ny), float64(dat[0])/(Nx*Ny)) // imprime as densidades
	dat[1] = 0 // zera as densidades para serem atualizadas a cada geração
	dat[0] = 0
	for t = 1; t <= NG; t++ { // conta as gerações
		for gd = 0; gd < Nx*Ny; gd++ { // conta todas as células da rede
			for i = 0; i < Nx; i++ {
				for j = 0; j < Ny; j++ {
					ind = j*Nx + i // célula principal; variáveis v abaixo dizem respeito à vizinhança de Moore
					vd = j*Nx + ((i + 1) % Nx)
					ve = j*Nx + ((i - 1 + Nx) % Nx)
					vb = ((j+1)%Ny)*Nx + i
					vc = ((j-1+Ny)%Ny)*Nx + i
					vd1 = ((j+1)%Ny)*Nx + (i+1)%Nx
					vd2 = ((j-1+Ny)%Ny)*Nx + (i-1+Nx)%Nx
					vd3 = ((j+1)%Ny)*Nx + (i-1+Nx)%Nx
					vd4 = ((j-1+Ny)%Ny)*Nx + (i+1)%Nx
					soma = phi[vd] + phi[ve] + phi[vb] + phi[vc] + phi[vd1] + phi[vd2] + phi[vd3] + phi[vd4] // soma dos estados da vizinhança
					if phi[ind] == 1 { // se a célula está viva
						if (soma == 2) || (soma == 3) { // se 2 ou 3 vizinhos também estão vivos
							updt[ind] = 1 // sobrevive
						} else {
							updt[ind] = 0 // passa para o estado morto

						}
					} else { // se a célula está morta
						if soma == 3 { // se há exatamente 3 vizinhos vivos
							updt[ind] = 1 // nasce
						} else {
							updt[ind] = 0 // permanece morta
						}
					}
				}
			}
		}
		for i = 0; i < Nx; i++ { // para todas as células da rede
			for j = 0; j < Ny; j++ {
				phi[j*Nx+i] = updt[j*Nx+i] // troca a rede inicial para a atualizada
				dat[phi[j*Nx+i]]++ // atualiza a densidade populacional
			}
		}
		fmt.Fprintf(f, "%.2f  %.2f\n", float64(dat[1])/(Nx*Ny), float64(dat[0])/(Nx*Ny)) // imprime a densidade populacional atualizada
		op(t, phi) // cria o arquivo .dat a cada geração (jdv-tcc-1, jdv-tcc-2...)
		fmt.Printf("%d/%d\n", t, NG) // mostra gerações completas em tempo real no terminal
		dat[1] = 0 // zera as densidades populacionais novamente
		dat[0] = 0
	}
}
