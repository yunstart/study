package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("study/mashibing/unit11/example/test.txt")
	if err != nil {
		fmt.Printf("文件打开出错,出错信息为%v\n", err)
	}
	fmt.Printf("文件=%v", file)
	err2 := file.Close()
	if err2 != nil {
		fmt.Printf("文件关闭出错，出错信息为%v\n", err2)
	}
}
