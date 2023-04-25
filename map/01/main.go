package main

import (
	"fmt"
)

func main() {
	var a map[string]int
	fmt.Println(a)
	b := make(map[string]int, 8)
	fmt.Println(b)
	b["abc"] = 100
	b["xyz"] = 200
	fmt.Println(b)
	fmt.Printf("type:%T\n", b)
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("b:%v\n", b)
	v, ok := b["abc"]
	fmt.Println(v, ok)
	//
	c := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("c:%#v\n", c)
	//iterate
	for _, v := range b {
		fmt.Println(v)
	}
	// sort iterate
	
	//delete
	delete(b, "abc")
	fmt.Println(b)
}
