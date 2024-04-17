#!/bin/bash
#
#SBATCH --mail-user=lamonts@cs.uchicago.edu
#SBATCH --mail-type=ALL
#SBATCH --job-name=ppsync 
#SBATCH --output=./slurm/out/%j.%N.stdout
#SBATCH --error=./slurm/out/%j.%N.stderr
#SBATCH --chdir=/home/lamonts/mpcs52060/ppsync/benchmark
#SBATCH --partition=debug 
#SBATCH --nodes=1
#SBATCH --ntasks=1
#SBATCH --cpus-per-task=12
#SBATCH --mem-per-cpu=900
#SBATCH --exclusive

alias go=/usr/lib/go-1.21/bin/go 

slurm_dir=./slurm/out/
locks=('tas' 'ttas' 'eb')
threads=('2' '4' '6' '8' '12')
increment=10000000
runs=10
file=data.txt 

if [[ -e "$file" ]]; then 
    rm -rf "$file"
fi 

for lock in "${locks[@]}"; do
    printf "# %s lock timings\n" "$lock" 
    printf "# %s lock timings\n" "$lock" >> "$file"
    for thread in "${threads[@]}"; do
        total=0 
        for ((i = 0 ; i < "$runs" ; i++)); do 
            if [[ "eb" == "$lock" ]]; then 
                got=$(go run ppsync/experiment $lock 32 1024 $increment $thread) 

            else
                got=$(go run ppsync/experiment $lock $increment $thread) 
            fi 
            total=$(bc -l <<< "scale=10;  $total+$got")
        done 
        total=$(bc -l <<< "scale=2;  $total/$runs")
        printf "%d %0.2f\n" $thread $total  
        printf "%d %0.2f\n" $thread $total >> "$file"
    done 
    printf "\n\n" >> "$file" 
done

gnuplot benchmark.gs
