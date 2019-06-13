package main

import "log"

func main() {
	// make a channel of ints with a capacity of 1
	ch := make(chan int, 1)

	// start a goroutine to listen on the channel
	go func() {
		for {
			log.Printf("popped: %d", <-ch)
		}
	}()

	// loop and spawn a goroutine to push i onto the channel
	for i := 0; i < 100; i++ {
		go func(n int) {
			log.Printf("pushing: %d", n)
			ch <- n
		}(i)
	}

	select {}
}
