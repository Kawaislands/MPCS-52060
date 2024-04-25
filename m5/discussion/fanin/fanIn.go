package main

import (
	"fmt"
	"math/rand"
	"time"
)

func counterService(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
func main() {
	c := fanIn(counterService("Bob"), counterService("Sally"))
	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Done.")
}
