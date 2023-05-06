package main

import "fmt"

//结构体的匿名字段,字段类型唯一
type Person struct {
	string
	int8
}

func main() {
	p1 := Person{
		"zhangsan",
		22,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
}

