package main

import "fmt"

func main() {
	//定义管道，声明一个int类型的管道
	var intChan chan int
	//make初始化
	intChan = make(chan int, 3)
	//证明管道是引用类型
	//intChan的值: 0xc000020100
	fmt.Printf("intChan的值: %v\n", intChan)
	//向管道存放数据
	intChan <- 10
	num := 20
	intChan <- num
	intChan <- 40
	//管道中读取数据
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Println("num1: ", num1)
	fmt.Println("num2: ", num2)
	fmt.Println("num3: ", num3)
	//注意：在没有使用协程的情况下，如果管道的数据已经全部取出，再取就会报错
	//输出管道的长度
	fmt.Printf("管道的实际长度:%v,管道的容量:%v", len(intChan), cap(intChan))
}
