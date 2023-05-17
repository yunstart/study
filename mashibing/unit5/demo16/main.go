package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//now函数返回值是结构体
	fmt.Printf("now的类型为%T,值为：%v\n", now, now)
	fmt.Println(now) //2023-05-05 11:24:44.8308035 +0800 CST m=+0.003920101
	//调用结构体方法，获取时，分，秒
	fmt.Printf("年：%v,月:%v,日:%v\n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("时：%v,分:%v,秒:%v\n", now.Hour(), now.Minute(), now.Second())
	//日期的格式化,直接输出
	fmt.Printf("当前年月日: %d-%d-%d 时分秒: %d:%d:%d\n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	//格式化日期，字符串变量接收
	datestr1 := fmt.Sprintf("当前年月日: %d-%d-%d 时分秒: %d:%d:%d\n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Println(datestr1)
	//函数格式化日期
	//这个参数字符串各个数字必须是固定的，必须这样写
	//2006年一月2号下午3点4分5秒
	datestr2 := now.Format("2006/01/02 15/04/05")
	fmt.Println(datestr2)
	datestr3 := now.Format("2006-01-02 15:04:05")
	fmt.Println(datestr3)
}
