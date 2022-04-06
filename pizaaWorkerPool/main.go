package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type customer struct {
	PizzaType  int
	reciptTime time.Time
}

type seller struct {
	job      customer
	duration time.Duration
}

var (
	size    = 5
	clients = make(chan customer, size)
	data    = make(chan seller, 10)
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		time.Sleep(time.Second)
		result := time.Now().Sub(c.reciptTime)
		output := seller{c, result}
		data <- output

	}
	w.Done()
}

func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := customer{i, time.Now()}
		clients <- c
	}
	close(clients)
}

func main() {
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of workers:", cap(data))

	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		os.Exit(1)
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	go create(nJobs)
	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Printf("Pizza: %d", d.job.PizzaType)
			fmt.Printf("reciept time %v", d.job.reciptTime)
			fmt.Printf("duration :%d", d.duration)
			fmt.Println("is ready")
		}
		finished <- true
	}()

	makeWP(nWorkers)
	fmt.Printf(": %v\n", <-finished)
}
