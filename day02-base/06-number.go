package main

import (
	"fmt"
	"reflect"
)
func main() {
	/**
	1、数字类型
	 */
	var a int8 = 4
	var b int32 = 4
	var c int64 = 4
	d := 4
	fmt.Printf("a: %T %v \n", a, a)
	fmt.Printf("b: %T %v \n", b, b)
	fmt.Printf("c: %T %v \n", c, c)
	fmt.Printf("d: %T %v \n", d, d)


	/**
	2、布尔值
	 */

	var isError = false
	fmt.Printf("isError: %T %v \n", isError, isError)

	// reflect.TypeOf查看数据类型
	fmt.Println(reflect.TypeOf(c))
}
/*
a: int8 4
b: int32 4
c: int64 4
d: int 4

int64
*/