/* MODELO RPS DE MAY-LEONARD
criado em       : 2022
ult. atualização: 2025/05/24
autor           : Beatriz Bittencourt <beatrizdecbittencourt@gmail.com>
notas           : Executa o modelo RPS de May-Leonard (mobilidade, reprodução, predação) em arquivos .dat.
compilação      : -
execução        : go run ml_rps.go
*/


package main
import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Nx, Ny, NG, NF, NS, min, max, pm, ppred, pr = 250, 250, 1000, 10, 3, 0, 4, 0.8, 0.1, 0.1
)

var (
	i, j, num, vizinho, iind, jind, indn, gd, t, l int
	a0, r float64
	dst [4]int
	phi[Nx*Ny]int
)

func ic(){
	for i := 0; i < Nx; i++ {
		for j := 0; j < Ny; j++ {
				phi[i*Ny+j] = rand.Intn(max-min) + min
		}
	}
}

func op(num int, phi [Nx*Ny]int){
	f := fmt.Sprintf("mlrps-%v.dat", num)
	file, _ := os.Create(f)
	for i := 0; i < Nx; i++ {
		for j := 0; j < Ny; j++ {
			fmt.Fprintf(file, "%d ", phi[i*Ny+j])
		}
		fmt.Fprintf(file, "\n")
	}
	defer file.Close()
}

func main() {
	r = float64(NG)/NF
	a0 = r
	pp := [NS*NS]int{0, 1, 0, 0, 0, 1, 1, 0, 0}
	rand.Seed(time.Now().UnixNano())
	ic()
	op(0, phi)
	
	for i := 0; i < Nx*Ny; i++ {
		dst[phi[i]]++;				
	}
	f, _ := os.Create("pop-mlrps.dat")
	for i := 1; i < 4; i++ {
		fmt.Fprintf(f, "%f ", float64(dst[i])/(Nx*Ny))
	}
	fmt.Fprintf(f, "%f\n", float64(dst[0])/(Nx*Ny));
	
	for t = 1; t <= NG; t++ {
		gd = 0;
		for gd < Nx*Ny {
			loop:
				iind := rand.Intn(Nx)
				jind := rand.Intn(Ny)
				indn := jind*Nx+iind
				if phi[indn] == 0 {
					goto loop
				}
			switch rand.Intn(4) {
				case 0:
					vizinho = jind*Nx+((iind+1)%Nx)
				case 1:
					vizinho = jind*Nx+((iind-1+Nx)%Nx)
				case 2:
					vizinho = ((jind+1)%Ny)*Nx+iind
				default:
					vizinho = ((jind-1+Ny)%Ny)*Nx+iind
			}
			at := rand.Float64()				
			if (at < pm) {
				phi[indn], phi[vizinho] = phi[vizinho], phi[indn]
				gd++
			} else {
				if (at < (pm + ppred)) {
					if phi[vizinho] != 0 && pp[(phi[indn]-1)*3+phi[vizinho]-1] == 1 {
						dst[phi[vizinho]]-- 
						phi[vizinho]= 0
						dst[0]++
						gd++
					}	
				} else {
					if phi[vizinho] == 0 {
						phi[vizinho] = phi[indn]
						dst[phi[vizinho]]++
						dst[0]--
						gd++
					}	
				}		
			}
		}
		for i := 1; i < 4; i++ {
			fmt.Fprintf(f, "%f ", float64(dst[i])/(Nx*Ny))
		}
		fmt.Fprintf(f, "%f\n", float64(dst[0])/(Nx*Ny));
	
		if(float64(t) >= a0){
			a0+= r
			l++
			op(l, phi)
			fmt.Printf("%d/%d\n", t, NG)
		}
	}
	defer f.Close()
}
