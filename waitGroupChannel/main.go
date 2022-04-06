package main

import (
	"fmt"
	"sync"
	"time"
)

type Material struct {
	material int
	mutex    sync.Mutex
}

func main() {
	rawMaterial := Material{material: 10000}

	var waitGroup sync.WaitGroup

	for i := 0; i < 100; i++ {
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
	time.Sleep(time.Second * 2)
	fmt.Printf("Pizza %d is ready \n", num)
}
