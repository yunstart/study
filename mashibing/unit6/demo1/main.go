package main

import "fmt"

func foo() {
	defer func() {
		//recover
		err := recover()
		if err != nil {
			fmt.Println("捕获到错误")
			fmt.Println("err: ", err)
		}
	}()
	n1 := 10
	n2 := 0
	ret := n1 / n2
	fmt.Println(ret)
}
func main() {
	foo()
	fmt.Println("code continue...")
}
