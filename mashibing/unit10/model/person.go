package model

import "fmt"

type person struct {
	Name string
	age  int
}

// 定义构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 定义set和get函数，对age字段进行封装，方法中可以对字段加一系列的限制确保字段的安全合理性
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("输入的年龄错误")
	}
}

func (p *person) GetAge() int {
	return p.age
}
