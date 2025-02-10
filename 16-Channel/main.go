package main

import (
	"fmt"
)

// func main() {
// 	ch := make(chan string) // 📡 Tạo một channel kiểu string

// 	go func() {
// 		ch <- "Hello from goroutine!" // 🚀 Gửi dữ liệu vào channel
// 	}()

// 	message := <-ch      // ⏳ Nhận dữ liệu từ channel
// 	fmt.Println(message) // ✅ In ra: Hello from goroutine!
// }

func main() {
	done := make(chan int, 10)
	//	mở	ra	N	goroutine
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("Hello	World")
			done <- 1
		}()
	}
	//	đợi	cả	10	goroutine	hoàn	thành
	for i := 0; i < cap(done); i++ {
		<-done
	}
}
