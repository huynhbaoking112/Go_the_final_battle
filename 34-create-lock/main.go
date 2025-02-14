package main

import (
	"fmt"
	"sync"
)

type Lock struct {
	lock chan struct{}
}

func getLock(l *Lock) bool {
	isLock := false
	select {
	case <-l.lock:
		isLock = true
	default:
	}
	return isLock
}

func UnLock(l *Lock) {
	l.lock <- struct{}{}
}

func NewLock() Lock {
	var l Lock
	l.lock = make(chan struct{}, 1)
	l.lock <- struct{}{}
	return l

}

func main() {

	l := NewLock()
	var wg sync.WaitGroup
	counter := 0
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if getLock(&l) {
					counter++
					UnLock(&l)
					fmt.Println("Counter is", counter)
					return
				}
			}
		}()
	}
	wg.Wait()
}
