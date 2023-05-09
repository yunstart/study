package main

import (
	"fmt"
	"sync"
)

// 定义一个变量
var totalNum int
var wg sync.WaitGroup

// 加入互斥锁
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		//加锁
		lock.Lock()
		totalNum = totalNum + 1
		//解锁
		lock.Unlock()
	}
}
func sub() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		totalNum = totalNum - 1
		lock.Unlock()
	}
}
func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(totalNum)
}
