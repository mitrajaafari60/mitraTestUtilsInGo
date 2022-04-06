package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("------- start number of 100 goroutin-----------")

	var waitGroup sync.WaitGroup

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Println("Preparing Pizza:", x)
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("finished")
}
