package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	minute := now.Minute()
	fmt.Println(year, month, day, minute)
	fmt.Println(now.Unix())
	n := 2
	time.Sleep(time.Duration(n) * time.Second)
	t1 := now.Add(time.Hour)
	fmt.Println(t1,t1.Location())
}
