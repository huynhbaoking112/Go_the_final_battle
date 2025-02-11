package main

import (
	"log"
	"net"
	"net/rpc"
)

// định	nghĩa	service	struct
type HelloService struct{}

// định	nghĩa	hàm	service	Hello,	quy	tắc:
// 1.	Hàm	service	phải	public	(viết	hoa)//	2.	Có	hai	tham	số	trong	hàm
// 3.	Tham	số	thứ	hai	phải	kiểu	con	trỏ
// 4.	Phải	trả	về	kiểu	error
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello	" + request
	//	trả	về	error	=	nil	nếu	thành	công
	return nil
}
func main() {
	// đăng ký một dịch vụ RPC với tên "HelloService" để cho phép client gọi từ xa

	// new(HelloService) Tạo một con trỏ đến một struct mới của HelloService tương đương với &HelloService{}

	// Điều này cần thiết vì trong Go, RPC yêu cầu service phải có các phương thức (method) có receiver là con trỏ (ví dụ: func (p *HelloService) Hello(...))

	rpc.RegisterName("HelloService", new(HelloService))
	//	chạy	rpc	server	trên	port	1234
	// Mở một socket TCP trên PORT 1234, để chương trình có thể lắng nghe (listen) các kết nối từ client
	listener, err := net.Listen("tcp", ":1234")
	//	nếu	có	lỗi	thì	in	ra
	if err != nil {
		log.Fatal("ListenTCP	error:", err)
	}
	//	vòng	lặp	để	phục	vụ	nhiều	client
	for {
		//	accept	một	connection	đến
		conn, err := listener.Accept()
		//	in	ra	lỗi	nếu	có
		if err != nil {
			log.Fatal("Accept	error:", err)
		}
		//	phục	vụ	client	trên	một	goroutine	khác
		//	để	giải	phóng	main	thread	tiếp	tục	vòng	lặp
		go rpc.ServeConn(conn)
	}
}
