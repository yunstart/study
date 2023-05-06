package main

import "fmt"

//定义结构体
type person struct {
	name, city string
	age        int8
}

func main() {
	var p person
	p.name = "zhangsan"
	p.city = "beijing"
	p.age = 20
	fmt.Println(p)
	fmt.Printf("p=%#v\n", p)
	//匿名结构体
	var user struct {
		name    string
		married bool
	}
	user.name = "zhangsan"
	user.married = false
	fmt.Printf("%#v", user)
}
