package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//写入文件操作
	//打开文件
	file, err := os.OpenFile("study/mashibing/unit11/example/demo.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败,err:\n", err)
	}
	//关闭文件
	defer file.Close()
	//写入文件操作：io流
	writer := bufio.NewWriter(file)
	// writer.WriteString("你好")
	for i := 0; i < 10; i++ {
		writer.WriteString("你好,Golang!\n")
	}
	writer.Flush()
}
