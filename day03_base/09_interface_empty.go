package main

import "fmt"

func show(a interface{}) {
	fmt.Printf("值：%v, 类型：%T\n", a, a)
}

func main() {

	/**
	空接口
		也可以直接当做类型来使用，可以表示任意类型
	*/
	// 空接口作为函数的参数
	show(1)
	show("123")
	show([3]int{3, 10, 12}) // 数组
	show([]int{3, 10, 12})  // slice
	show(map[string]string{})

	// 切片实现空接口  类似于python中的list，可以放任何类型
	result := []interface{}{"张三", 22, 3000, true}
	fmt.Println(result) // [张三 22 3000 true]

	// map 的值实现空接口
	studentInfo := make(map[string]interface{})
	studentInfo["name"] = "李四"
	studentInfo["age"] = 33
	studentInfo["salary"] = 8000
	studentInfo["isMale"] = true
	fmt.Println(studentInfo) // map[age:33 isMale:true name:李四 salary:8000]

	/**
	空接口，类型断言
		一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。
		要判断空接口中值的类型：  x.(T)
	*/

	var x interface{}
	x = "hello world"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("非字符串类型")
	}

}
