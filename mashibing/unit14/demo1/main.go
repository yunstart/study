package main

import (
	"fmt"
	"reflect"
)

// 利用一个函数，函数的参数定义为一个空接口
func testReflect(i interface{}) {
	//调用TypeOf函数，返回reflect.Type类型数据
	reType := reflect.TypeOf(i)
	fmt.Println("reType:", reType)
	fmt.Printf("reType的具体类型:%T", reType)
	//调用ValueOf函数，返回reflect.Value类型数据
	reValue := reflect.ValueOf(i)
	fmt.Println("reValue:", reValue)
	fmt.Printf("reValue的具体类型:%T", reValue)
	//调用ValueOf函数，返回reflect.Value类型数据
	//如要获得reValue的数值，要调用Int()方法，返回v持有的有符号整数
	num2 := 100 + reValue.Int()
	fmt.Println(num2)
	//reValue转成空接口
	i2 := reValue.Interface()
	//类型断言
	n := i2.(int)
	n2 := n + 30
	fmt.Println(n2)
}
func main() {
	fmt.Println("")
	//对基本数据类型进行反射
	//定义一个基本数据类型
	var num int = 100
	testReflect(num)

}
