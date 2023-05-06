package main

import "fmt"

//结构体的构造函数
//调用一个函数得到一个结构体的实例
//定义结构体
type person struct {
	name, city string
	age        int8
}

// func newPerson(name, city string, age int8) person {
// 	return person{
// 		name: name,
// 		city: city,
// 		age:  age,
// 	}
// }
//结构体是值拷贝，使用指针提高性能
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
func main() {
	p1 := newPerson("zhangsan", "beijing", 22)
	fmt.Println(p1)
}
