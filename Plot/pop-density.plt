set terminal pdfcairo
set output "pop-density.pdf"

set decimalsign ','
set xlabel offset 0.0, 0.0 "{/Times t}"
set ylabel offset 0.0, 0.0 "{/Times {Density (total = 1)}" enhanced
set ytics offset 0.0, 0.0 
set xtics offset 0.0, 0.0 
set key at graph 0.5, 0.8 reverse Left samplen 1 horizontal center
set key noautotitle

set xrange [0:500]
set yrange [0:1]

plot \
"data1.dat" u ($0):($2) w l dt "-" lw 2.0 lc rgb "red" enhanced, \
"data1.dat" u ($0):($1) w l lw 2.0 lc rgb "red" t"{/Times {Title 1}" enhanced, \
"data2.dat" u ($0):($2) w l dt "-" lw 2.0 lc rgb "yellow"enhanced, \
"data2.dat" u ($0):($1) w l lw 2.0 lc rgb "yellow" t"{/Times {Title 2}" enhanced, \
"data3.dat" u ($0):($2) w l dt "-" lw 2.0 lc rgb "green" enhanced, \
"data3.dat" u ($0):($1) w l lw 2.0 lc rgb "green" t"{/Times {Title 3}" enhanced, \
"data4.dat" u ($0):($2) w l dt "-" lw 2.0 lc rgb "blue" enhanced, \
"data4.dat" u ($0):($1) w l lw 2.0 lc rgb "blue" t"{/Times {Title 4}" enhanced

unset output
