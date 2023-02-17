package main

import "fmt"

func main(){

	/**
	Array   [əˈreɪ]
	介绍：
		数组是指一系列 同一类型数据的集合, int, string。
		一个数组包含的元素个数被称为数组的长度
		数组是一个长度固定的数据类型,   [5]int 和 [10]int 是两个不同的类型
	特点：
		占用内存的连续性，而索引数组元素的速度非常快

	对立面： slice 切片  [slaɪs]
	 		可以增长和收缩的动态序列
	 */

	/**
	定义 + 赋值 = 初始化
	 */
	// 定义
	var a [5]int
	var b [3]string
	var c [2]bool

	fmt.Println(a)   // [0 0 0 0 0]
	fmt.Println(b)   // [  ]
	fmt.Println(c)   // [false false]

	// 赋值
	b[0] = "hello"
	b[1] = "world"
	fmt.Println(b)  // [hello world ]
	fmt.Println(len(b))  //3
	fmt.Println(b[2])

	//初始化
	// 数组长度不确定，可以使用 ... 代替数组的长度
	d := [3]int{100,99,98}
	e := [...]string{"a", "b", "c"}
	fmt.Println(d)
	fmt.Println(e)


	/**
	遍历
	 */
	var cityArray = [...]string{"beijing", "shanghai", "shenzhen"}
	// 普通遍历
	for i:=0; i< len(cityArray); i++{
		fmt.Println(i, cityArray[i])
	}

	// k,v遍历数组
	for index, value := range cityArray{
		fmt.Println(index, value)
	}
	// 0 beijing
	// 1 shanghai
	// 2 shenzhen

	for i:=0 ; i<len(b); i++{
		fmt.Println(i,b[i])
	}
	for index, value := range b{
		fmt.Println(index, value)
	}
	// 0 hello
	// 1 world
	// 2


	/**
	查找：在array中，查找某个值
	 */
	var numArray = [...]int{1, 2,3}
	fmt.Println(numArray)
	target := 2
	flag := false
	for i:=0; i< len(numArray); i ++{
		if numArray[i] == target{
			flag = true
			break
		}
	}
	fmt.Println(flag)


	/**
	修改：将usa修改为china
	 */
	countryArray := [...]string{"uk", "usa", "jp"}
	oldS := "usa"
	newS  := "china"
	fmt.Println(countryArray)

	var resultArray  [len(countryArray)] string
	for i:=0; i< len(countryArray); i++{
		if countryArray[i] == oldS{
			resultArray[i] = newS
			continue
		}
		resultArray[i] = countryArray[i]
	}
	fmt.Println(resultArray)  // [uk china jp]

	/**
	删除：删除目标元素
	 */
	foods := [3]string{"noodles", "rice", "meat"}
	targetFood := "rice"

	var newFoods [len(foods)] string
	for i:=0; i< len(foods); i++{
		if foods[i] == targetFood{
			continue
		}
		newFoods[i] = foods[i]
	}
	fmt.Println(newFoods)  // [noodles  meat]
	fmt.Println(newFoods[0])  // noodles
	fmt.Println(newFoods[1])  //
	fmt.Println(newFoods[2])  // meat

}
