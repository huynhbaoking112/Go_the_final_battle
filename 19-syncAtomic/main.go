package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var value int64

func work(m *sync.WaitGroup) {

	defer m.Done()
	var i int64
	for i = 0; i <= 100; i++ {
		atomic.AddInt64(&value, i)
	}

}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	work(&wg)
	work(&wg)

	wg.Wait()

	fmt.Println("ket thuc", value)

}
