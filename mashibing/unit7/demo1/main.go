package main

import (
	"fmt"
)

func main() {
	var scores [5]int
	scores[0] = 95
	scores[1] = 91
	scores[2] = 35
	scores[3] = 65
	scores[4] = 25
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	fmt.Println("sum:", sum)
}
