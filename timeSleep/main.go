package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("------- start number of 100 goroutin-----------")

	for i := 0; i < 100; i++ {

		go func(x int) {

			fmt.Println("Preparing Pizza:", x)
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("finished")
}
