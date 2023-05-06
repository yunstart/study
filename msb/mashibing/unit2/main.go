package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	str := fmt.Sprintf("Name= %v, Age= %v", s.Name, s.Age)
	return str
}
func main() {
	//方法1：
	var s1 Student = Student{"张三", 20}
	fmt.Println(s1)
	//方法2：
	var s2 Student = Student{
		Name: "张三",
		Age:  22,
	}
	fmt.Println(s2)
	//方法3：返回结构体的指针类型
	var s3 *Student = &Student{"张三", 20}
	fmt.Println(s3)
	var s4 *Student = &Student{
		Name: "李四",
		Age: 20,
	}
	fmt.Printf("s4 value is: %#v\n",s4)
	fmt.Printf("s4 type is: %T", s4)
}
