package main

import (
	"fmt"
	"reflect"
)

// 利用一个函数，函数的参数定义为空接口
func testReflect(i interface{}) {
	//1.调用TypeOf函数，返回reflect.TypeOf类型数据
	reType := reflect.TypeOf(i)
	//2.调用ValueOf函数，返回reflect.ValueOf类型数据
	reValue := reflect.ValueOf(i)
	//获取变量的类别
	//1.reType
	k1 := reType.Kind()
	fmt.Println(k1)
	//2.reValue
	k2 := reValue.Kind()
	fmt.Println(k2)

	//获取变量的类型
	i1 := reValue.Interface()
	n, flag := i1.(Student)
	if flag {
		fmt.Printf("结构体的类型是 %T", n)
	}
}

// 定义学生结构体
type Student struct {
	Name string
	Age  int
}

func main() {
	//对结构体类型进行反射
	//定义结构体的具体实例
	s := Student{
		Name: "张三",
		Age:  22,
	}
	testReflect(s)
}
