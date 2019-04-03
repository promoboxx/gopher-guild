package main

import (
	"log"
	"math/rand"
	"sync"
)

// Expectation:
// The for loop will push 10 random ints onto the
// channel then the read function will pop each of
// them off the channel and print them out
func main() {
	// unbuffered channel
	ch := make(chan int)
	// wait groups are used to ensure that something
	// completely finishes its work before exiting the scope
	wg := &sync.WaitGroup{}

	// start a go routine that has a receiover for the channel
	go read(ch, wg)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// push a random number onto the channel
		n := rand.Intn(1000)
		ch <- n
		log.Printf("pushed: %d", n)
	}

	wg.Wait()
}

func read(ch chan int, wg *sync.WaitGroup) {
	// pop from the channel
	for {
		n := <-ch
		log.Printf("popped: %d", n)
		wg.Done()
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
