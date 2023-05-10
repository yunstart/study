package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		//创建一个切片，将读取的数据放入切片
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		//将读取内容在服务器端输出
		fmt.Println(string(buf[0:n]))
	}
}

func main() {
	fmt.Println("服务器启动了...")
	//进行监听，需要指定TCP协议
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败, err", err)
		return
	}
	//监听成功后，等待客户端连接
	for {
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("客户端等待失败,err2:", err2)
		} else {
			//连接成功
			fmt.Printf("等待连接成功,conn=%v,接收的客户端信息是:%v\n", conn, conn.RemoteAddr().String())
		}
		//准备一个协程，协程处理客户端服务请求
		go process(conn) //不同的客户端请求，连接Conn不一样
	}
}
