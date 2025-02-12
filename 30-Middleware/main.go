package main

import (
	"fmt"
	"net/http"
	"time"
)

// hàm	business	logic	của	chúng	ta
func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

// đây	là	một	hàm	thực	hiện	việc	đóng	gói	hàm	truyền	vào
// để	ghi	nhận	thời	gian	thực	thi	service
// hàm	này	được	xem	như	là	một	Middleware
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		//	ghi	nhận	thời	gian	trước	khi	chạy
		timeStart := time.Now()
		//	next	là	hàm	business	logic	được	truyền	vào
		next.ServeHTTP(wr, r)
		//	tính	toán	thời	gian	thực	thi
		timeElapsed := time.Since(timeStart)
		//	log	ra	thời	gian	thực	thi
		fmt.Println(timeElapsed)
	})
}

// Trong	đó,	http.Handler:
//
//	type	Handler	interface	{
//				ServeHTTP(ResponseWriter,	*Request)
//	}
func main() {
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
	err := http.ListenAndServe(":8080", nil)
	//	...
	if err != nil {
		panic("Error hehe")
	}

}
