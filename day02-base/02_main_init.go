package main

import "fmt"

var x int8 = 10
const PI = 3.14
const E = 2.17

/**
init函数和main函数
相同点：
	两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用。
不同点：
	init可以应用于任意包中，且可以重复定义多个。
	main函数只能用于main包中，且只能定义一个

执行顺序：
	对同一个go文件的 init() 调用顺序是从上到下的
 */
func main()  {
	fmt.Println(PI)

}


func init()  {
	fmt.Println(E)
}


func init()  {
	fmt.Println(x)
}

/**
2.17
10
3.14
 */
