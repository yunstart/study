package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func writeData(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("写入的数据为: ", i)
		time.Sleep(time.Second)
	}
	//管道关闭
	close(intChan)
}
func readData(intChan chan int) {
	defer wg.Done()
	//遍历管道
	for v := range intChan {
		fmt.Println("读取的数据为: ", v)
		time.Sleep(time.Second)
	}
}
func main() {
	//写协程和读协程共同操作一个管道
	//定义管道
	intChan := make(chan int, 50)
	//开启读和写的协程
	wg.Add(2)
	go writeData(intChan)
	go readData(intChan)
	wg.Wait()
}
