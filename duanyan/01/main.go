package main

import "fmt"

func main() {
	var x interface{}
	x = 500
	fmt.Println(x)
	s, ok := x.(string)
	if !ok {
		fmt.Println(s)
	}
	switch t := x.(type) {
	case string:
		fmt.Println(t)
	case bool:
		fmt.Println(t)
	default:
		fmt.Println(t)
	}
}
