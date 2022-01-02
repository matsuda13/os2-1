set title "relation of buffersize and processing speed"
set xlabel "buffersize"
set ylabel "processing speed(Mb/s)"
set grid

set output "out.svg"
plot "./out/graphdata5000000.dat" with linespoints pointtype 7
</svg>