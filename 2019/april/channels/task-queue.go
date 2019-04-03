package main

import (
	"log"
	"math/rand"
	"time"
)

type task struct {
	t    string
	work func() int
}

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

func main() {
	tasks := make(chan task)

	go runTasks(tasks)

	go createTasks(tasks)

	select {}
}
