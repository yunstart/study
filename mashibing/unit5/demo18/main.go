package main

import "fmt"

func main() {
	num1 := new(int)
	var num2 *int = new(int)
	fmt.Println(num1, num2)
	arr1 := new([3]int)
	fmt.Println(arr1)
	type student *struct {
		name string
		age  int
	}
	var s1 *student = new(student)
	fmt.Println(s1)
}
