package main

import (
	"fmt"
	"study/mashibing/unit10/model"
)

func main() {
	p := model.NewPerson("张三")
	p.SetAge(22)
	fmt.Println(p.GetAge())
	fmt.Println(*p)
}
