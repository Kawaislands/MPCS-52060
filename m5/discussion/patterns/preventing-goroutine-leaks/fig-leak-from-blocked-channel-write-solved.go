package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	newRandStream := func(done <-chan interface{}, wg *sync.WaitGroup) <-chan int {
		randStream := make(chan int)
		go func() {
			defer wg.Done()
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()

		return randStream
	}
	var wg sync.WaitGroup
	wg.Add(1)
	done := make(chan interface{})
	randStream := newRandStream(done, &wg)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	// Simulate ongoing work
	wg.Wait()
}
