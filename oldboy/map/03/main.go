package main

import "fmt"

func main() {
	//定义切片，切片内的元素类型为map
	var mapSlice = make([]map[string]int, 8, 8)

	fmt.Println(mapSlice)
	// mapSlice[0] = make(map[string]int, 8)
	// mapSlice[0]["abc"] = 123
	// mapSlice[1] = make(map[string]int, 8)
	// mapSlice[1]["xyz"] = 456
	//切片内的map元素需要make初始化
	for i := 0; i < len(mapSlice); i++ {
		mapSlice[i] = make(map[string]int, len(mapSlice))
	}
	mapSlice[0]["abc"] = 123
	for _, v := range mapSlice {
		fmt.Println(v)
	}
	// fmt.Println(mapSlice[1])
}
