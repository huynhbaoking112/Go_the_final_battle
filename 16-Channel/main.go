package main

import (
	"fmt"
)

func main() {
	ch := make(chan string) // 📡 Tạo một channel kiểu string

	go func() {
		ch <- "Hello from goroutine!" // 🚀 Gửi dữ liệu vào channel
	}()

	message := <-ch      // ⏳ Nhận dữ liệu từ channel
	fmt.Println(message) // ✅ In ra: Hello from goroutine!
}
