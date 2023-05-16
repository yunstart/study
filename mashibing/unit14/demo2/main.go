package main

import (
	"fmt"
	"reflect"
)

// 利用一个函数，函数的参数定义为空接口
func testReflect(i interface{}) {
	reType := reflect.TypeOf(i)
	fmt.Println("reType:", reType)
	reValue := reflect.ValueOf(i)
	fmt.Println("reValue:", reValue)

	i2 := reValue.Interface()
	n, flag := i2.(Student)
	if flag {
		fmt.Println("学生的姓名：", n.Name)
		fmt.Println("学生的年龄：", n.Age)
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
