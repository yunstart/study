package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//读取文件
	//返回内容为：[]byte,err
	//不需要进行open\close操作，文件的打开和关闭操作被封装在ReadFile函数内部
	content, err := ioutil.ReadFile("study/mashibing/unit11/example/test.txt")
	if err != nil {
		fmt.Printf("读取出错，错误为%v\n", err)
	} else {
		//读取成功，内容显示在终端
		fmt.Printf("文件内容:%v\n", content)
		fmt.Printf("文件内容:%v\n", string(content))
	}

}
