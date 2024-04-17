package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func increment_blocking(counter *int32, atomicValue *int32, group *sync.WaitGroup) {

	for !atomic.CompareAndSwapInt32(atomicValue, 0, 1) {
		runtime.Gosched() //allow other goroutines to do stuff.
	}

	// ---------------------------- Start of the Critical Section
	*counter = *counter + 1

	// ---------------------------- Start of the Critical Section
	atomic.StoreInt32(atomicValue, 0)
	group.Done()
}

func increment_nonblocking(counter *int32, atomicValue *int32, group *sync.WaitGroup) {

	//Lock Free
	localCounter := *counter

	for !atomic.CompareAndSwapInt32(counter, localCounter, localCounter+1) {
		localCounter = *counter
	}

	group.Done()
}
func main() {

	var blockCounter int32
	var nonblockCounter int32
	var atomicValue int32
	var group sync.WaitGroup

	group.Add(1000)
	for i := 0; i < 1000; i++ {
		go increment_blocking(&blockCounter, &atomicValue, &group)
	}

	group.Add(1000)
	for i := 0; i < 1000; i++ {
		go increment_nonblocking(&nonblockCounter, &atomicValue, &group)
	}

	group.Wait()
	fmt.Printf("Block Counter = %v\n", blockCounter)
	fmt.Printf("NonBlock Counter = %v\n", nonblockCounter)

}
