package main

import "fmt"

type cInterface interface {
	c()
}
type bInterface interface {
	b()
}
type aInterface interface {
	bInterface
	cInterface
	a()
}
type student struct {
}

func (s student) a() {
	fmt.Println("a")
}
func (s student) b() {
	fmt.Println("b")
}
func (s student) c() {
	fmt.Println("c")
}

func main() {
	var s student
	var a aInterface = s
	a.a()
	a.b()
	a.c()
}
