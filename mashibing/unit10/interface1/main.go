package main

import "fmt"

//接口定义规则、规范、某种能力
type SayHello interface {
	//声明没有实现的方法
	SayHello()
}

//接口的实现，定义一个结构体
//中国人结构体
type Chinese struct {
	Name string
}

//实现接口的方法
func (person Chinese) SayHello() {
	fmt.Println("你好")
}

//接口的实现，定义一个结构体
//美国人结构体
type American struct {
	Name string
}

//实现接口的方法
func (person American) SayHello() {
	fmt.Println("Hello")
}

//定义一个函数，接收具备SayHello接口的能力的变量
func greet(s SayHello) {
	s.SayHello()
}

//自定义数据类型
type integer int

func (i integer) SayHello() {
	fmt.Println("say hi", i)
}
func main() {
	var i integer = 10
	var s SayHello = i
	s.SayHello()
	//Chinese实例
	c := Chinese{}
	//American实例
	a := American{}
	//Chinese方法
	greet(c)
	//American方法
	greet(a)
	var arr [3]SayHello
	arr[0] = Chinese{"张三"}
	arr[1] = American{"jack"}
	arr[1] = American{"tom"}
	fmt.Println(arr)

	// var s SayHello = c
	// s.SayHello()
}
