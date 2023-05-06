package main

import "fmt"

//嵌套结构体字段冲突
type Address struct {
	Province   string
	City       string
	UpdateTime string
}
type Email struct {
	Addr       string
	UpdateTime string
}
type Person struct {
	Name   string
	Gender string
	Age    int
	Address
	Email
}

func main() {
	p1 := Person{
		Name:   "zhangsan",
		Gender: "male",
		Age:    22,
		Address: Address{
			Province:   "shandong",
			City:       "qingdao",
			UpdateTime: "2022-01-01",
		},
		Email: Email{
			Addr:       "zhangsan@local.com",
			UpdateTime: "2022-07-01",
		},
	}
	fmt.Println(p1.Province)
	fmt.Println(p1.Addr)
	fmt.Println(p1.Address.UpdateTime)
	fmt.Println(p1.Email.UpdateTime)
}
