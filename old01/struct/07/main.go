package main

import "fmt"

type Address struct {
	Province string
	City     string
}
type Person struct {
	Name    string
	Gender  string
	Age     int
	Address Address //嵌套另一个结构体
}

func main() {
	p1 := Person{
		Name:   "zhangsan",
		Gender: "male",
		Age:    22,
		Address: Address{
			Province: "shandong",
			City:     "qingdao",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.Address.Province)
}
