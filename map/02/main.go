package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var scoreMap = make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)
	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }
	keys := make([]string, 0, len(scoreMap))
	for k := range scoreMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
