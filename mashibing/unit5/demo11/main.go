package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}
	result := add(1, 2)
	fmt.Println(result)
}
