package main

import "fmt"

func getSum() func(int) int {
	var sum int=0
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

func main() {
	a := getSum()
	// x := a(10)
	fmt.Println(a(1))
	fmt.Println(a(2))
	fmt.Println(a(3))
}
