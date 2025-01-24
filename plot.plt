Nx=150
Ny=150
NF=500
set terminal png size Nx+55, Ny+55 crop
ext="png"
unset xtics
unset ytics
unset colorbox

set xrange  [0:Nx-1]
set yrange  [0:Ny-1]
set border
set size ratio -1
unset key

set palette defined ( 0 '#ffffff',\
                      1 '#790cff',\
                      2 '#3176f8',\
                      3 '#18cee3',\
                      4 '#6dfebd',\
                      5 '#d0e183',\
                      6 '#fd954c',\
					  7 '#fe1a0d')

i= 0
while (i <= NF ){
	set output sprintf("png-%d.%s", i, ext)
	plot sprintf("data-%d.dat", i) \
	     u ($1+1):($2+1):($3) matrix w image
	i= i+ 1
}
unset output
