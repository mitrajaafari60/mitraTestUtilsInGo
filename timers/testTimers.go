package main

import (
	"fmt"
	"time"
)

func main() {
	certainSomething := true // will cause time loop to repeat

	timeDelay := 900 * time.Millisecond // == 900000 * time.Microsecond

	var endTime <-chan time.Time // signal for when timer us up

	for {
		// if a certain something happens, start a timer
		if certainSomething && endTime == nil {
			endTime = time.After(timeDelay)
		}
		select {
		case <-endTime:
			fmt.Println("Yes Finally!")
			endTime = nil
		default:
			fmt.Println("not yet")
			time.Sleep(50 * time.Millisecond) // simulate work
			continue
		}

		// function shouldn't be called if the elapsed amount
		//      of time required has not yet been reached
		iShouldOnlyBeCalledWhenDelayHasBeenReached() // this could also just be moved to the <- endtime block above
	}
}

func iShouldOnlyBeCalledWhenDelayHasBeenReached() {
	fmt.Println("I've been called")
}
