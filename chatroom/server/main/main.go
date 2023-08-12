package main

import (
	_ "errors"
	"fmt"
	"net"
	"time"

	"../model"
)

func init() {
	initPool("47.100.60.204:6379", 16, 0, 300*time.Second)
	initUserDao()
}

func processMain(conn net.Conn) {
	//读取客户端发送的信息
	ps := &Processor{
		Conn: conn,
	}

	err := ps.process()
	if err != nil {
		fmt.Println("client connecting err=", err)
		return
	}

}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	fmt.Println("服务器[2020]监听8889端口")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	for {
		fmt.Println("等待客户端来链接服务器......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		defer conn.Close()
		//一旦链接成，则启动一个协程和客户端保持通信
		go processMain(conn)
	}
}
