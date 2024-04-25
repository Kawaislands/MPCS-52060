// pi estimates pi using goroutines and monte carlo method
//
//Usage: monte_carlo <interval> <threads>
//    interval = the number of iterations to perform
//    threads = the number of threads (i.e., goroutines to spawn)
//
//  Author:, Lamont Samuels
//  Date: 11/08/18
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const usage = "Usage: pi <interval> <threads>\n" +
	"\tinterval = the number of iterations to perform" +
	"\tthreads = the number of worker threads"

type goContext struct {
	workChan chan int
	total    int
	workers  int
	rGen     *rand.Rand
}

const mainThreadRank = 0

// computeCircles computes the number of circles given an upper-bound nad lower-bound
func computeCircles(start, end int, rGen *rand.Rand) int {

	var localCircles int
	var x, y, distanceSquared float64

	for iter := start; iter < end; iter++ {
		x = rGen.Float64()
		y = rGen.Float64()
		distanceSquared = x*x + y*y
		if distanceSquared <= 1 {
			localCircles++
		}
	}
	return localCircles
}

// performWork is the function that is called to perform the intervals of work
func performWork(ctx *goContext, rank, localInterval int) {

	//Define local interval
	start := localInterval * rank
	end := start + localInterval

	localCircles := computeCircles(start, end, ctx.rGen)

	//Check to see if the thread is either the main thread
	if rank == mainThreadRank {
		ctx.total += localCircles
		for j := 0; j < ctx.workers; j++ {
			localCircles = <-ctx.workChan
			ctx.total += localCircles
		}
	} else {
		//Compute the local number of circles and send it to the workChannel
		ctx.workChan <- computeCircles(start, end, ctx.rGen)
	}

}

func main() {

	//Check to make sure we received the correct number of arguments (not required by assignment)
	if len(os.Args) < 3 {
		panic(fmt.Errorf(usage))
	}

	//Retrieve the command-line arguments and perform conversion if needed
	threadCount, _ := strconv.Atoi(os.Args[2]) // + 1 //Add 1 for the main thread
	intervals, _ := strconv.Atoi(os.Args[1])

	//Initilize the waitgroup and shared goroutine contexts
	s1 := rand.NewSource(time.Now().UnixNano())
	rGen := rand.New(s1)
	context := goContext{rGen: rGen, workChan: make(chan int), workers: threadCount - 1}

	//Start the gorountines
	i := mainThreadRank + 1
	localInterval := intervals / threadCount

	for ; i < threadCount; i++ {
		go performWork(&context, i, localInterval)
	}
	//Make the main thread perform some work
	performWork(&context, mainThreadRank, localInterval)

	//Print out the estimate
	piEstimate := 4.0 * float64(context.total) / float64(intervals)

	fmt.Printf("Estimated pi: %e\n", piEstimate)
}
