package main

import (
	"log"
	"math/rand"
)

// Expectation:
// To be able to create a buffered channel size 10
// push 10 random numbers onto the channel
// pop the 10 random numbers off the channel
func main() {
	// channels can act as queues and will maintain order
	// based on when something was inserted
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		n := rand.Intn(1000)
		log.Printf("pushed: %d", n)
		ch <- n
	}

	for len(ch) > 0 {
		log.Printf("popped: %d", <-ch)
	}
}

// Result:
// Experienced results as expected
