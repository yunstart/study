package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type Student struct {
	Name string
	Age  int
}

// 结构体方法
func (s Student) Print() {
	fmt.Println("调用了Print()方法")
	fmt.Println("学生的名字是: ", s.Name)
}
func (s Student) GetSum(n1, n2 int) int {
	fmt.Println("调用了GetSUM方法")
	return n1 + n2
}
func (s Student) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

// 定义函数操作结构体
func TestStudentStruct(a interface{}) {
	//将a转为reflect.Value类型
	val := reflect.ValueOf(a)
	fmt.Println(val)
	//通过reflect.Value类型操作结构体内部字段
	n1 := val.NumField()
	fmt.Println(n1)
	//遍历获取具体的字段
	for i := 0; i < n1; i++ {
		fmt.Printf("第%d个字段的值为: %v\n", i, val.Field(i))
	}
	//通过reflect.Value类型操作结构体内部方法
	n2 := val.NumMethod()
	fmt.Println(n2)
	//调用print方法
	//调用方法，方法的首字母必须大写才能有对应的反射的访问权限
	val.Method(1).Call(nil)
	//调用GetSum方法
	//定义Value的切片
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	result := val.Method(0).Call(params)
	fmt.Println("GetSum的返回值:", result[0].Int())
}
func main() {
	s := Student{
		Name: "张三",
		Age:  12,
	}
	TestStudentStruct(s)
}
