package main

import (
	"fmt"
	"sync"
)

// Lock	try	lock
type Lock struct {
	c chan struct{}
}

// tạo	một	lock
func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

// Lock	try	lock,	trả	về	kết	qủa	lock	là	true/false
func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

// Unlock	,	giải	phóng	lock
func (l Lock) Unlock() {
	l.c <- struct{}{}
}

// biến	đếm	toàn	cục
var counter int

func main() {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if l.Lock() {
					counter++
					fmt.Println("current	counter", counter)
					l.Unlock()
					return
				}
			}
			// if !l.Lock() {
			// 	//	log	error
			// 	// fmt.Println("lock failed")
			// } else {
			// 	counter++
			// 	fmt.Println("current	counter", counter)
			// 	l.Unlock()
			// }
		}()
	}
	wg.Wait()
}
