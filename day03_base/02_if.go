package main

import "fmt"

func main()  {

	/**
	常规写法
	 */
	score := 85
	res := "A"
	if score >90{
		res = "A"
	}else if score > 60{
		res = "B"
	}else{
		res = "C"
	}
	fmt.Println("你的KPI是：", res)


	/**
	特殊写法
	 */
	var level string  = "专家"

	//这里的 score 是局部作用域
	if salary:= 20; salary >50{
		level = "专家"
	}else if salary > 20{
		level = "资深"
	}else{
		level = "高级"
		// 只能在函数内部打印 score
		fmt.Println("你的职级是：", level)
	}
	fmt.Println("你的职级是：", level)

	//fmt.Println(salary)  // fmt.Println(score) //undefined: score
}