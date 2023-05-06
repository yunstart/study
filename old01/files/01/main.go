package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("err:%#v\n", err)
		return
	}
	defer fileObj.Close()
	var tmp = make([]byte, 128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Printf("err:%#v\n", err)
		return
	}
	fmt.Println(n)
	fmt.Println(string(tmp))
}
