package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	value int
}

func work(so int, m *sync.WaitGroup) {

	defer m.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		fmt.Println("Them tu ", so)
		total.Unlock()
	}

}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	work(1, &wg)
	work(2, &wg)

	wg.Wait()

	fmt.Println("ket thuc", total.value)

}
