package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	start := time.Now()
	fmt.Println(start)
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(t, elapsed)

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(4 * time.Second)
}
