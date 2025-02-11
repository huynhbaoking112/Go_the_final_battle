package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			return
		}
	}
}
func main() {
	//	nhận	vào	context	parent	(Background)	và	trả	về	context	child	(ctx)	và	hàm	cancel
	//	deadline	10	secs
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	//	đơn	giản	hoá	worker	pool	cho	ngắn	gọn
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}
	time.Sleep(time.Second)
	//	mặc	dù	ctx	sẽ	expire	theo	timeout	đã	set	trước	đó
	//	ta	vẫn	gọi	cancel	để	đóng	context	child	và	các	children	của	nó
	//	để	tránh	giữ	chúng	tồn	tại	không	cần	thiết
	cancel()
	//	sử	dụng	waitGroup	thay	cho	done	channel
	wg.Wait()
}

// func worker(king chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		default:
// 			fmt.Println("hello")
// 		case <-king:
// 			return
// 		}
// 	}
// }
// func main() {
// 	//	nhận	vào	context	parent	(Background)	và	trả	về	context	child	(ctx)	và	hàm	cancel
// 	//	deadline	10	secs

// 	king := make(chan bool)

// 	var wg sync.WaitGroup
// 	//	đơn	giản	hoá	worker	pool	cho	ngắn	gọn
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go worker(king, &wg)
// 	}
// 	time.Sleep(time.Second)
// 	//	mặc	dù	ctx	sẽ	expire	theo	timeout	đã	set	trước	đó
// 	//	ta	vẫn	gọi	cancel	để	đóng	context	child	và	các	children	của	nó
// 	//	để	tránh	giữ	chúng	tồn	tại	không	cần	thiết
// 	close(king)
// 	//	sử	dụng	waitGroup	thay	cho	done	channel
// 	wg.Wait()
// }
