package main

import (
	"net"
	"net/rpc"
	"time"
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
	rpc.RegisterName("HelloService", new(HelloService))
	for {
		//	chủ	động	gọi	tới	client
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		rpc.ServeConn(conn)
		conn.Close()
	}
}
