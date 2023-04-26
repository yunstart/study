package main

import "fmt"

func foo(name string) func() {
	// name = "func()"
	return func() {
		fmt.Println("Hello", name)
	}

}
func main() {
	//闭包
	// var name string
	// fmt.Println("Please enter your full name: ")
	// fmt.Scanln(&name)
	ret := foo("func()")
	ret()
}
