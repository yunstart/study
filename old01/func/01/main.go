package main

import "fmt"

func sayHello(name string) {
	fmt.Println("hello", name)
}

// func intSum(a, b int) int {
// 	ret := a + b
// 	return ret
// }
func intSum(a, b int) (ret int) {
	ret = a + b
	return ret
}

func intSum2(a ...int) int {
	// fmt.Printf("%T\n", a)
	ret := 0
	for _, arg := range a {
		ret += arg
	}
	return ret
}
func intSum3(a int, b ...int) int {
	ret := a
	for _, arg := range b {
		ret = ret + arg
	}
	return ret
}
func main() {
	name := "azc"
	sayHello(name)
	ret := intSum(2, 3)
	fmt.Println(ret)
	ret2 := intSum2(2, 3, 5)
	fmt.Println(ret2)
	ret3 := intSum3(2, 3, 5)
	fmt.Println(ret3)
}
