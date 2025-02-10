package main

import (
	"sync/atomic"
	"time"
)

func loadConfig() map[string]string {
	return make(map[string]string)
}
func requests() chan int {
	return make(chan int)
}
func main() {
	//	nắm	giữ	thông	tin	cấu	hình	của	server
	var config atomic.Value
	//	khởi	tạo	giá	trị	ban	đầu
	config.Store(loadConfig())
	go func() {
		//	cập	nhật	thông	tin	sau	mỗi	10	giây
		for {
			time.Sleep(10 * time.Second)
			config.Store(loadConfig())
		}
	}()
	//	tạo	nhiều	worker	sử	lý	request
	//	dùng	thông	tin	cấu	hình	gần	nhất
	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				c := config.Load()
				//	xử	lý	request	với	cấu	hình	c
				_, _ = r, c
			}
		}()
	}
}
