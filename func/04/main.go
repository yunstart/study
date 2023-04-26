package main

import "fmt"

func main() {
	//匿名函数
	sayHi := func() {
		fmt.Println("anonymous func")
	}
	sayHi()
}
