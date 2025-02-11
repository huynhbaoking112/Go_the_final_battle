package main

import (
	"fmt"
	"time"
)

// Producer: Tạo ra một chuỗi số nguyên dựa trên bội số factor và gửi vào channel
func Producer(factor int, out chan<- int) { // out là send-only
	for i := 0; ; i++ {
		out <- i * factor
		fmt.Println("Producer")
	}
}

// Consumer: Nhận các số từ channel và in ra màn hình
func Consumer(in <-chan int) { // in là receive-only
	for v := range in {
		fmt.Println("Consumer", v)

	}
}

func main() {
	// Hàng đợi có buffer 64 phần tử
	ch := make(chan int, 64)

	// Tạo chuỗi số với bội số 3
	go Producer(3, ch)

	// Tạo chuỗi số với bội số 5
	go Producer(5, ch)

	// Tạo consumer
	go Consumer(ch)

	// Để tránh main() kết thúc sớm, ta có thể thêm time.Sleep() hoặc WaitGroup
	time.Sleep(5 * time.Second) // Chặn main() không cho kết thúc
}
