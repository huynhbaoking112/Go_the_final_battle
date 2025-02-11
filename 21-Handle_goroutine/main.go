package main

import (
	"fmt"
	"time"
)

// Mục đích của pool worker - Xử lí nhiều tác vụ với một số hữu hạn goroutine nhất định, tránh trường hợp tạo hàng triệu goroutine để xử lí hàng triệu request

// Ham xu li
func worker(q chan int, workerId int, r chan bool) {
	for i := range q {
		// gia lap thoi gian xu li
		time.Sleep(time.Second)
		fmt.Println("Xu li tu ", workerId, "task", i)
		r <- true
	}
}

func main() {

	start := time.Now()

	q := make(chan int)

	r := make(chan bool)

	//Tao goroutine xu li
	maxGoroutine := 9

	for i := 0; i < maxGoroutine; i++ {
		go worker(q, i, r)
	}

	// Giao task
	maxTask := 17
	for i := 1; i <= maxTask; i++ {
		go func(i int) {
			q <- i
		}(i)
	}

	for i := 1; i <= maxTask; i++ {
		<-r
	}

	fmt.Println("Thoi gian chay:", time.Since(start))

}
