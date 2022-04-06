package main

import (
	"fmt"
	"sync"
	"time"
)

type Material struct {
	material int
	mutex    sync.RWMutex
}

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

func read() int {
	return <-readValue
}

func backing() {
	var value int
	var internalCounter int
	for {
		select {
		case newValue := <-writeValue:
			internalCounter = internalCounter + 1
			value = newValue
			putInOven(value)

		case readValue <- value:
			internalCounter = internalCounter - 1
		}
	}
}
func main() {
	rawMaterial := Material{material: 10000}

	var waitGroup sync.WaitGroup

	for i := 1; i <= 100; i++ {
		waitGroup.Add(1)
		go perparePizza(&waitGroup, &rawMaterial, i)

	}
	for i := 0; i < 10; i++ {
		go backing()
	}
	waitGroup.Wait()
	fmt.Println("finished with remained material:", rawMaterial.material)
}

func perparePizza(wg *sync.WaitGroup, m *Material, num int) {
	defer wg.Done()
	//	fmt.Println("Preparing Pizza:", num)
	m.mutex.Lock()
	m.material--
	m.mutex.Unlock()
	//	time.Sleep(time.Second * 2)
	set(num)

	fmt.Printf("Pizza %d is ready \n", num)
}

func putInOven(num int) {
	//	fmt.Println("putInOven", num)
	time.Sleep(time.Second * 3)
}
