package main

import "fmt"

type cat struct{}

func (c cat) say() {
	fmt.Println("miao,miao~")
}

type dog struct{}

func (d dog) say() {
	fmt.Println("wang,wang~")
}

type sayer interface {
	say()
}

func da(arg sayer) {
	arg.say()
}
func main() {
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)

}
