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
