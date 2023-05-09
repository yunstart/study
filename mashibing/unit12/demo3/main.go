package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //只定义无需赋值
func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1) //协程开始的时候加1操作
		go func(n int) {
			defer wg.Done() //协程执行完成减1
			fmt.Println(n)
		}(i)
	}
	//主线程一直在阻塞，当wg减为0，就停止
	wg.Wait()
}
