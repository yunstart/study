package main
//new and make
import "fmt"

func main() {
	var a *int = new(int)
	*a = 100
	fmt.Println(*a)

	var b map[string]int = make(map[string]int, 5)
	b["zhangsan"] = 100
	fmt.Println(b)
}
