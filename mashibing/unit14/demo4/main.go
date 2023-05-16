package main

import (
	"fmt"
	"reflect"
)

// 利用一个函数，函数的参数定义为空接口
func testReflect(i interface{}) {
	reValue := reflect.ValueOf(i)
	reValue.Elem().SetInt(30)
}

func main() {
	var num int = 100
	testReflect(&num)
	fmt.Println(num)
}
