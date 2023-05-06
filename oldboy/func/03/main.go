package main

import "fmt"

func add(x, y int) int {
	return x + y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func main() {
	z := calc(2, 3, add)
	fmt.Println(z)
}
