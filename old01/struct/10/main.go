package main

import (
	"encoding/json"
	"fmt"
)

// 字段可见性
type Student struct {
	ID   int
	Name string
}

func newStudent(id int, name string) Student {
	return Student{
		ID:   id,
		Name: name,
	}
}

type Class struct {
	Title string `json:"title"`
	//Students []Student `json:"student_list" db:"student" xml:"ss"`
	Students []Student 
}

func main() {
	c1 := Class{
		Title:    "ban1",
		Students: make([]Student, 0, 20),
	}
	for i := 0; i < 10; i++ {
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("%#v\n", c1)
	//json序列化
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed,err:", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)
	jsonstr := `{"Title":"ban1","Students":[{"ID":0,"Name":"stu00"},{"ID":1,"Name":"stu01"},{"ID":2,"Name":"stu02"},{"ID":3,"Name":"stu03"}]}`
	var c2 Class
	err = json.Unmarshal([]byte(jsonstr), &c2)
	if err != nil {
		fmt.Println("json unmarshal failed", err)
		return
	}
	fmt.Println(c2)
}
