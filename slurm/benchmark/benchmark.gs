set terminal pngcairo  # Output format
set output 'lock.png'  # Output filename
set title "Execution time for (1e7 iterations)"
set xlabel "Threads"
set ylabel "Execution Time(seconds)"


plot 'data.txt' index 0 using 1:2 with lines title 'tas lock',\
     'data.txt' index 1 using 1:2 with lines title 'ttas lock',\
     'data.txt' index 2 using 1:2 with lines title 'eb lock'