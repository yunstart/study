package main

import "fmt"

func main() {
	var arr [3]int8
	fmt.Println(len(arr))
	fmt.Printf("arr: %p\n", &arr)
	fmt.Printf("arr: %p\n", &arr[0])
	fmt.Printf("arr: %p\n", &arr[1])
	fmt.Printf("arr: %p\n", &arr[2])
}
