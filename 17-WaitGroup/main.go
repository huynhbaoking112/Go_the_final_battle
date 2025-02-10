package main

import (
	"fmt"
	"sync"
	"time"
)

func download(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Đảm bảo gọi Done() dù có lỗi
	fmt.Println("Downloading:", url)
	time.Sleep(2 * time.Second) // Giả lập tải chậm
	fmt.Println("Completed:", url)
}

func main() {
	var wg sync.WaitGroup
	urls := []string{"https://golang.org", "https://google.com", "https://github.com"}

	for _, url := range urls {
		wg.Add(1)             // Tăng số lượng goroutines cần chờ
		go download(url, &wg) // Chạy goroutine tải URL
	}

	wg.Wait() // Chờ tất cả tải xong
	fmt.Println("All downloads finished!")
}
