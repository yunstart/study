package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("study/mashibing/unit11/example/test.txt")
	if err != nil {
		fmt.Println("文件打开失败,err=", err)
	}
	//当函数退出时，让file关闭，防止内存泄露
	defer file.Close()
	//创建文件流
	reader := bufio.NewReader(file)
	//读取操作
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(str)
			break
		} else {
			fmt.Println(str)
		}
	}
	//end
	fmt.Println("文件读取成功并完毕")
}
