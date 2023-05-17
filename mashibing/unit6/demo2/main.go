package main

import (
	"errors"
	"fmt"
)

func foo() (err error) {
	n1 := 10
	n2 := 0
	if n2 == 0 {
		return errors.New("除数不能为0")
	} else {
		result := n1 / n2
		fmt.Println(result)
		return nil
	}
}
func main() {
	err := foo()
	if err != nil {
		fmt.Println("自定义错误:", err)
		panic(err)
	}
}