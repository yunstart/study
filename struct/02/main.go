package main

import "fmt"

//结构体指针
type person struct {
	name, city string
	age        int8
}

func main() {
	//结构体实例化
	var p2 = new(person)
	fmt.Printf("%T,%p\n", p2,&p2)
	fmt.Println(*p2)

	// (*p2).name = "zhangsan"
	// (*p2).city = "beijing"
	// (*p2).age = 20
	p2.name = "zhangsan"
	p2.city = "beijing"
	p2.age = 20
	fmt.Println(p2)
	fmt.Printf("%#v\n", p2)

	//取结构体的地址进行实例化
	// p3 := person{}
	// fmt.Printf("%T\n", p3)
	// fmt.Printf("%#v\n", p3)
	// p3.name = "zhangsan"
	// p3.city = "beijing"
	// p3.age = 20
	// fmt.Printf("%#v", p3)
	fmt.Printf("person{}:%T\n", person{})
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("%#v\n", p3)
	p3.name = "zhangsan"
	p3.city = "beijing"
	p3.age = 20
	fmt.Printf("%#v", p3)
}
