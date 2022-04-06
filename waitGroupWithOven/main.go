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

var (
	mOven  sync.Mutex
	inOven int
)

func main() {
	rawMaterial := Material{material: 10000}

	var waitGroup sync.WaitGroup

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go perparePizza(&waitGroup, &rawMaterial, i)

	}
	waitGroup.Wait()
	fmt.Println("finished with remained material:", rawMaterial.material)
}

func perparePizza(wg *sync.WaitGroup, m *Material, num int) {
	defer wg.Done()
	fmt.Println("Preparing Pizza:", num)
	m.mutex.Lock()
	m.material--
	m.mutex.Unlock()
	//	time.Sleep(time.Second * 2)
	var isCooking chan bool
	for {
		ovenManeger(num, isCooking)
		select {
		case <-isCooking:
			putInOven(num)
		default:
			fmt.Println("waiting for accepting cook!", num)
			time.Sleep(time.Second)
		}
	}

	fmt.Printf("Pizza %d is ready \n", num)
}

func ovenManeger(num int, couldPleaseCook chan bool) {
	if inOven < 10 {
		mOven.Lock()
		inOven = inOven + 1
		couldPleaseCook <- true
		mOven.Unlock()
	}
}
func putInOven(num int) {
	fmt.Println("putInOven", num)
	time.Sleep(time.Second * 30)
	mOven.Lock()
	inOven = inOven - 1
	mOven.Unlock()
}
