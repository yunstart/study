package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	//关闭管道
	close(intChan)
	// intChan <- 30
	n1 := <-intChan
	fmt.Println(n1)
}
