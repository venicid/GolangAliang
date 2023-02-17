package main

import "fmt"



func main(){

	/**
	1、变量

	var 变量名 类型 = 表达式
	 */
	var name string = "zhang"
	var age int = 33

	var isRich bool
	var height, weight float32
	var (
		a string
		b float64
		c bool
	)

	a,b,c = "test", 33.333, false

	isRich = false
	height = 180.88
	weight = 180.88

	// 类型推导方式定义变量
	salary := 1888.3

	/*
	2、  fmt包
	 */

	// 判断数据的类型
	// 1. Printf 是格式化输出
	// https://www.runoob.com/go/go-fmt-sprintf.html
	fmt.Printf("%T\n", age)

	// Print  中间没有空格 不会自动换行
	fmt.Print(name,age)
	fmt.Print("\n")

	// 3. Println 中间有空格 会自动换行
	fmt.Println(name, age)

	fmt.Println(salary)
	fmt.Println(isRich)
	fmt.Println(height, weight)
	fmt.Println(a,b,c)


	/**
	3、常量
	 */
	const PI = 3.14
	const E = 2.7182
	fmt.Println(PI, E)
}
