package main

import (
	"fmt"
	"math/rand"
	"time"
)

func counterService(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func main() {
	c := make(chan string)
	go counterService("Count", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
	}
	fmt.Println("Done.")
}
