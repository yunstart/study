package main

import "fmt"

//方法和接收者
//定义结构体
type Person struct {
	name string
	age  int8
}

//Person类型的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//为Person类型定义方法
func (p Person) Dream() {
	fmt.Printf("%s \n", p.name)
}

//定义修改age的方法
//接收者的类型是指针
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}
func main() {
	p1 := NewPerson("zhangsan", 22)
	(*p1).Dream()
	p1.Dream()
	p1.SetAge(33)
	fmt.Printf("%v", p1)
}
