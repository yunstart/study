package main

import "fmt"

//getSum参数为空
//返回值为一个函数，函数的参数是int类型的参数，返回值也是int类型
func getSum() func(int) int {
	var sum int = 0
	return func(i int) int {
		sum += i
		return sum
	}
}
func main() {
	f := getSum()
	fmt.Println(f(1))
	fmt.Println(f(2))
}
