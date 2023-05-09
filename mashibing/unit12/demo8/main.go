package main

import "fmt"

func main() {
	//定义管道//初始化管道
	var intChan = make(chan int, 100)
	// intChan = make(chan int, 100)
	//管道赋值
	for i, count := 0, 10; i < count; i++ {
		intChan <- i
	}
	//关闭管道
	close(intChan)
	//遍历前，如果没有关闭管道，会出现deadlock
	//fatal error: all goroutines are asleep - deadlock!
	//遍历管道: for-range

	for v := range intChan {
		fmt.Println("value: ", v)
	}

}
