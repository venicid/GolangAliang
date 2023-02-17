package main

import "fmt"

func main()  {

	/**
	关于指针：只需要记住两个符号：&（取地址）和 *（根据地址取值）
	 */
	var a = 10

	fmt.Printf("%d\n", a)  // 10
	fmt.Printf("%T\n", a)  // int

	fmt.Printf("%d\n", &a)  // &a 指针地址 (824633761976)
	fmt.Printf("%T\n", &a)   //  %T 指针类型 (*int )

	fmt.Printf("%d\n", *&a)  // *&a 指针取值 (10)

	var b *int = &a
	fmt.Printf("%d\n", *b)  // 10   指针的取值
	fmt.Printf("%T\n", b)  // *int  	指针的类型
	fmt.Printf("%d\n", b)  // 824633761944  指针的值
	fmt.Printf("%d\n", &b)  // 824634605600  指针的地址


	/**
	& 取变量的地址
	 */
	var m  = 10
	var n = &m
	var x = *n

	fmt.Println(m) // 10 a的值
	fmt.Println(n) // 0xc00001e060 a变量的内存地址
	fmt.Println(x) // 10 *内存地址 取值
}