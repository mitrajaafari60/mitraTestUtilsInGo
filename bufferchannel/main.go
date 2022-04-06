package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(c chan int) {
	stopReading := false
	i := 0
	for {
		select {
		case input := <-c:
			i++
			fmt.Println(i, input, len(c))
		default:
			c = nil
			fmt.Println("finished reading ", i)
			stopReading = true
		}

		if stopReading {
			break
		}
	}
}

func send(c chan int) {
	stopSending := false
	for {
		select {
		case c <- rand.Intn(10):
		default:
			fmt.Println("Not enough space")
			stopSending = true
		}
		if stopSending {
			break
		}
	}
}

func main() {
	c := make(chan int, 100)
	go add(c)

	//	time.Sleep(3 * time.Second)
	fmt.Println("start sending")
	go send(c)

	time.Sleep(3 * time.Second)

}
