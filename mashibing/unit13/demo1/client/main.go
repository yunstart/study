package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端启动...")
	//调用Dial函数,参数需要指定tcp协议，需要指定服务器端的IP+Port
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败:err", err)
		return
	}
	fmt.Println("客户端连接成功,conn:", conn)
	//通过客户端发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin) //os.stdin代表终端输入
	//从终端读取一行用户输入的信息
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("终端输入失败,err:", err)
	}
	//将str发送给服务器
	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("连接失败,err:", err)
	}
	fmt.Printf("终端数据通过客户端发送成功，一共发送了%d字节的数据，并退出", n)
}
