package main

import (
	"fmt"
	"sync"
)

type Runnable func(arg interface{})

type SharedContext struct {
	mutex *sync.Mutex
	group *sync.WaitGroup
	count int
}

func Increment(arg interface{}) {

	//Convert the generic context argument into its actual type by using type assertion
	var ctx *SharedContext
	ctx = arg.(*SharedContext)

	ctx.mutex.Lock()
	ctx.count += 1
	ctx.mutex.Unlock()
}

func ExecuteRunnable(runnable Runnable, arg interface{}) {

	go func() {
		runnable(arg)
		var ctx *SharedContext
		ctx = arg.(*SharedContext)
		ctx.group.Done()
	}()
}

func main() {

	var numOfThreads int
	numOfThreads = 3

	var runnables []Runnable
	runnables = make([]Runnable, numOfThreads)

	var ctx *SharedContext
	ctx = &SharedContext{&sync.Mutex{}, &sync.WaitGroup{}, 0}
	ctx.group.Add(numOfThreads)

	for i := 0; i < numOfThreads; i += 1 {
		runnables[i] = Increment
	}

	for i := 0; i < numOfThreads; i += 1 {
		go ExecuteRunnable(runnables[i], ctx)
	}
	ctx.group.Wait()
	fmt.Printf("Count=%d\n", ctx.count)

}
