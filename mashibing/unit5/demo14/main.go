package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	str := "golang你好"
	fmt.Println([]rune(str))
	fmt.Println([]byte(str))
	fmt.Println(strconv.Itoa(20320))
	fmt.Println(unicode.Han)
}
