package main

import "fmt"

//结构体初始化
//定义结构体
type person struct {
	name, city string
	age        int8
}

func main() {
	//键值对初始化
	p4 := person{
		name: "zhangsan",
		city: "beijing",
		age:  22,
	}
	fmt.Printf("%#v\n", p4)

	p5 := &person{
		name: "zhangsan",
		city: "beijing",
		age:  22,
	}
	fmt.Printf("%#v\n", p5)

	//值的列表进行初始化
	p6 := person{
		"zhangsan",
		"beijing",
		22,
	}
	fmt.Printf("%#v\n", p6)
}
