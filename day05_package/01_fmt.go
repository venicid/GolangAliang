package main

import "fmt"

func main() {
	/**
	print
		%v 按值的本来值输出
		%#v 输出 Go 语言语法格式的值
		%T 输出 Go 语言语法格式的类型和值
		%d 整型以十进制方式显示
	*/
	fmt.Print("zhangsan", "lisi", "wangwu")   // zhangsanlisiwangwu
	fmt.Println("zhangsan", "lisi", "wangwu") // zhangsan lisi wangwu
	name := "zhangsan"
	age := 20
	fmt.Printf("%s 今年 %d 岁\n", name, age)       // zhangsan 今年 20 岁
	fmt.Printf("值：%v --> 类型: %T\n", name, name) // 值：zhangsan --> 类型: string

	/**
	Sprint：字符串，填充拼接
	*/
	s1 := "sudo su %s;"
	s11 := fmt.Sprintf(s1, "deploy") // 字符串填充
	s2 := "hostname;"
	s3 := "whoami;"
	s4 := fmt.Sprintln(s11, s2, s3) // 字符串拼接
	fmt.Println(s4)                 // sudo su deploy; hostname; whoami;

}
