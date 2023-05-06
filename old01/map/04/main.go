package main

import "fmt"

//定义值为切片的map
func main() {
	var sliceMap = make(map[string][]int, 8)
	fmt.Println(sliceMap)
	_, ok := sliceMap["china"]
	if ok {
		fmt.Println(sliceMap["china"])
	} else {
		//切片初始化
		sliceMap["china"] = make([]int, 8)
		sliceMap["china"][0] = 100
		sliceMap["china"][1] = 200
	}
	fmt.Println(sliceMap)
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
}
