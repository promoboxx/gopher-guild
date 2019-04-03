package main

import (
	"log"
	"math/rand"
	"time"
)

// type that will handle some form of work
type task struct {
	t    string
	work func() int
}

// a function that loops infinitely and pushes a random task onto the task channel
func createTasks(tasks chan task) {
	possibleTasks := []task{
		{t: "logger", work: func() int { log.Printf("this is a logging task"); return 0 }},
		{t: "logger", work: func() int { log.Printf("this task is being worked on"); return 1 }},
		{t: "sleeper", work: func() int { log.Printf("sleeping for 5 seconds"); time.Sleep(5 * time.Second); return 1 }},
		{t: "adder", work: func() int { log.Printf("adding two random numbers"); return rand.Intn(100) + rand.Intn(100) }},
	}

	t := time.NewTicker(time.Second)

	for {
		select {
		case <-t.C:
			tasks <- possibleTasks[rand.Intn(len(possibleTasks)-1)]
		}
	}
}

// waits until a task is received on the task queue
// the calls its work function and does the work
func runTasks(tasks chan task) {
	for {
		select {
		case t := <-tasks:
			log.Printf("received task type: %s", t.t)
			res := t.work()
			log.Printf("exited: %d", res)
		}
	}
}

// An application that creates tasks at an interval with a ticker
// then reads them off the tasks channel and does that work
func main() {
	// make the channel
	tasks := make(chan task)

	// start listening for tasks on the channel
	go runTasks(tasks)

	// create the tasks
	go createTasks(tasks)

	// wait infinitely
	select {}
}
