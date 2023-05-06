package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "how do you do !"
	var wordCount = make(map[string]int, 10)
	// var wordCount = make(map[string]int, 10)
	words := strings.Split(s, " ")
	// fmt.Printf("%T", s1)
	for _, word := range words {
		fmt.Println(word)
		_, ok := wordCount[word]
		if ok {
			wordCount[word] += 1
		} else {
			wordCount[word] = 1
		}
	}
	fmt.Println(wordCount)
}
