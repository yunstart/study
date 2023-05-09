package main

import "fmt"

type aInterface interface {
	a()
}
type bInterface interface {
	b()
}
type student struct {
}

func (s student) a() {
	fmt.Println("aaa")
}
func (s student) b() {
	fmt.Println("bbb")
}
func main() {
	var s student
	var a aInterface = s
	var b bInterface = s
	a.a()
	b.b()
}
