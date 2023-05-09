package main

import "fmt"

type SayHello interface {
	SayHello()
}
type Chinese struct {
	Name string
}
type American struct {
	Name string
}

func (person Chinese) SayHello() {
	fmt.Println("你好")
}
func (person Chinese) dance() {
	fmt.Println("chinaese dance")
}
func (person American) SayHello() {
	fmt.Println("hi")
}
func (person American) sing() {
	fmt.Println("american sing")
}
func greet(s SayHello) {
	s.SayHello()
	//断言
	//看s是否能转成Chinese类型，并赋给ch变量
	// ch, ok := s.(Chinese)
	// if ok {
	// 	ch.dance()
	// } else {
	// 	fmt.Println("only chinese")
	// }
	//优化
	// if ch, ok := s.(Chinese); ok {
	// 	ch.dance()
	// } else {
	// 	fmt.Println("only chinese")
	// }
	switch s.(type) { //type属于go中的一个关键字，固定写法
	case Chinese:
		ch := s.(Chinese)
		ch.dance()
	case American:
		am := s.(American)
		am.sing()
	}
}

func main() {
	c := Chinese{
		Name: "张三",
	}
	a := American{
		Name: "jack",
	}
	greet(a)
	greet(c)
}
