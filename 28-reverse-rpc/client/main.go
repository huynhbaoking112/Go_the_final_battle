package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func doClientWork(clientChan <-chan *rpc.Client) {
	//		nhận	vào	đối	tượng	RPC	client	từ	channel
	client := <-clientChan
	//	đóng	kết	nối	với	client	trước	khi	hàm	exit
	defer client.Close()
	var reply string
	//	thực	hiện	lời	gọi	rpc	bình	thường
	err := client.Call("HelloService.Hello", "hello", &reply)
	fmt.Println(reply)
	if err != nil {
		fmt.Println("asd")
		log.Fatal(err)
	}
}

func main() {
	//	listen	trên	port 	1234	chờ	server	gọi
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP	error:", err)
	}
	clientChan := make(chan *rpc.Client)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accept	error:", err)
			}
			//	khi	mỗi	đường	link	được	thiết	lập,	đối	tượng
			//	RPC	client	được	khởi	tạo	dựa	trên	link	đó	và
			//	gửi	tới	client	channel
			clientChan <- rpc.NewClient(conn)
		}
	}()
	doClientWork(clientChan)
}
