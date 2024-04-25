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

func main() {
	bob := counterService("Bob's Count")
	sally := counterService("Sally's Count")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-bob) // Receive expression is just a value.
		fmt.Printf("You say: %q\n", <-sally)
	}
	fmt.Println("Done.")
}
