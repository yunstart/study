package main

import "fmt"

//指针

func main() {
	a := 10
	b := &a
	c := *b
	fmt.Println(*b, b, &b)
	fmt.Println(c, &c)
	fmt.Printf("%v,%p\n", a, &a)
}
