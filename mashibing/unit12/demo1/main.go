package main

import (
	"fmt"
	"strconv"
	"time"
)

// 协程
func foo() {
	for i := 0; i <= 1000; i++ {
		fmt.Println("Hello Golang + " + strconv.Itoa((i)))
		//阻塞1秒
		time.Sleep(time.Second)
	}
}
func main() { //主线程
	//开启协程
	go foo()
	for i := 0; i <= 10; i++ {
		fmt.Println("Hello msb + " + strconv.Itoa((i)))
		//阻塞1秒
		time.Sleep(time.Second)
	}
}
