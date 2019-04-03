package main

import (
	"log"
	"math/rand"
)

// Expectation:
// The for loop will push 10 random ints onto the
// channel then the read function will pop each of
// them off the channel and print them out
func main() {
	// unbuffered channel
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		// push a random number onto the channel
		n := rand.Intn(1000)
		ch <- n
	}

	read(ch)
}

func read(ch chan int) {
	// pop from the channel
	for len(ch) > 0 {
		n := <-ch
		log.Printf("popped: %d", n)
	}
}

// Result:
// This does not work because the unbuffered channel acts
// as though it is a full channel. Since there is nothing
// that is reading from the channel this causes the application
// deadlock
//
// From the godocs:
// If the channel is unbuffered, the sender blocks until the receiver has received the value.
// If the channel has a buffer, the sender blocks only until the value has been copied to the
// buffer; if the buffer is full, this means waiting until some receiver has retrieved a value.
// https://golang.org/doc/effective_go.html#channels
