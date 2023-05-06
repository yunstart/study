package main

import "fmt"

//自定义类型type
type MyInt int

//类型别名
// type byte=uint8
// type rune=int32
type NewInt = int

func main() {
	var a MyInt
	fmt.Printf("value:%v,type:%T\n", a, a)

	var b NewInt
	fmt.Printf("value:%v,type:%T\n", b, b)
}
