package main

import "fmt"

func main() {
	// var a []int
	// var b []string
	// fmt.Println(a, b)
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:4]
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
