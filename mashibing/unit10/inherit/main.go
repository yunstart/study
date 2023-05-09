package main

import "fmt"

type Animal struct {
	Age    int
	weight float32
}

func (an *Animal) Shout() {
	fmt.Println("大声喊叫")
}
func (an *Animal) showInfo() {
	fmt.Printf("动物的年龄是:%v,动物的体重是:%v\n", an.Age, an.weight)
}

type Cat struct {
	//复用，体现继承思想，嵌入匿名结构体
	Animal
}

func (c *Cat) scratch() {
	fmt.Println("挠人")
}
func main() {
	cat := &Cat{}
	cat.Age = 3
	cat.weight = 10.6
	cat.Shout()
	cat.showInfo()
	cat.scratch()
}
