package main

import (
	"log"
	"time"
)

const (
	signalOne = iota
	signalTwo
	signalThree
)

// Create 3 tickers that each go off at different points in time
// after each ticker goes off then push the appropriate signal
// on to the signal receiver. Once the signal receiver receives a
// signal then print out the signal that was received
func main() {
	// each ticker actually uses channels to alert that
	// it has gone off. Once the time has elapsed it pushes
	// the time that ticker went off onto its channel and then
	// that time can be read by a receiver
	ticker1 := time.NewTicker(time.Second * 5)
	ticker2 := time.NewTicker(time.Second * 10)
	ticker3 := time.NewTicker(time.Second * 15)
	signalReceiver := make(chan int)

	// setup the receiver first since we have an unbuffered channel
	go func(sr chan int) {
		// read forever on the signal receiver
		for {
			select {
			case sig := <-sr:
				log.Printf("received signal: %d", sig)
			}
		}
	}(signalReceiver)

	// start pushing data onto the signal reciever
	go func(t1, t2, t3 *time.Ticker, sr chan int) {
		for {
			select {
			case <-t1.C:
				sr <- signalOne
			case <-t2.C:
				sr <- signalTwo
			case <-t3.C:
				sr <- signalThree
			}
		}
	}(ticker1, ticker2, ticker3, signalReceiver)

	// wait forever
	select {}
}
