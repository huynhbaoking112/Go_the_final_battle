package main

import "fmt"

//	trả	về	channel	tạo	ra	chuỗi	số:	2,	3,	4,	...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

//	bộ	lọc:	xóa	các	số	có	thể	chia	hết	cho	số	nguyên	tố
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	//	chuỗi	số:	2,	3,	4,	...
	ch := GenerateNatural()
	for i := 0; i < 100; i++ {
		//	số	nguyên	tố	mới
		prime := <-ch
		//	Bộ	lọc	dựa	trên	số	nguyên	tố	mới
		fmt.Printf("%v:	%v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
		//	dựa	trên	chuỗi	số	còn	lại	trong	channel	để	lọc
		//	các	số	nguyên	tố	tiếp	theo	với	các	số	được
		//	trích	xuất	dưới	dạng	filter.	Các	channel	tương	ứng
		//	với	các	sàng	số	nguyên	tố	khác	nhau	được	kết	nối	liên	tiếp	nhau.
	}
}
