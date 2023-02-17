package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Getwd())
	os.Mkdir("go_demo", 0777)
	os.Rename("go_demo", "go_demo_new")
	os.Create("file.txt")

	// 打开文件并写入文件
	/*
	   O_RDONLY 打开只读文件
	   O_WRONLY 打开只写文件
	   O_RDWR 打开既可以读取又可以写入文件
	   O_APPEND 写入文件时将数据追加到文件尾部
	   O_CREATE 如果文件不存在，则创建一个新的文件
	*/
	file, _ := os.OpenFile("file.txt", os.O_RDWR|os.O_APPEND, 06666)
	_, err := file.WriteString("hello world")
	fmt.Println(err)

}
