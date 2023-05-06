package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s move...\n", a.name)
}

type Dog struct {
	Feet int8
	*Animal
}

func (d *Dog) wang() {
	fmt.Println("wangwangwang")
}
func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "lele",
		},
	}
	d1.move()
	d1.wang()
}
