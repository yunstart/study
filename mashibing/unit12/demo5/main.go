package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个变量
var totalNum int
var wg sync.WaitGroup

// 加入读写锁
var lock sync.RWMutex

func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println("开始读数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	lock.RUnlock()
}
func write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second * 10)
	fmt.Println("修改数据成功")
	lock.Unlock()
}
func main() {
	wg.Add(6)
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()
	wg.Wait()
	fmt.Println(totalNum)
}
