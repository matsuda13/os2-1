set terminal svg
set output "out.svg"
set title "relation of buffersize and processing speed"
set xlabel "buffersize(byte)"
set ylabel "processing speed(Mb/s)"
set grid
plot "./out/graphdata5000000.dat" with linespoints pointtype 7
</svg>