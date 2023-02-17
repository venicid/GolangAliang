package main

import "fmt"

func main(){

	/**
	定义并初始化
	 */

	var userInfo = map[string]string{
		"name": "alex",
		"passwd":"alex",
	}
	fmt.Println(userInfo)   // map[name:alex passwd:alex]


	/**
	增加
	 */
	cityMap := make(map[string]string)
	cityMap["beijing"] = "北京"
	cityMap["shanghai"] = "上海"
	fmt.Println(cityMap)  // map[beijing:北京 shanghai:上海]


	/**
	查找，
	 */
	 userInfo1 := map[string]string{
		"name": "alex",
		"passwd":"alex",
	}

	v, ok :=userInfo1["username"]
	if ok{
		fmt.Println(v)  // alex
	}else{
		fmt.Println("map中没有该元素")
	}

	 item := userInfo1["passwd"]
	 fmt.Print(item)  // 不存在


	 /**
	 删除
	  */
	 delete(userInfo1, "passwd")
	 fmt.Print(userInfo1)  // map[name:alex]


	/**
	遍历
	 */
	// 不支持普通遍历
	//for i:=0; i<len(userInfo1); i++{
	//	fmt.Println(i)
	//}
	scoreMap := map[string]int{
		"zhangsan": 24,
		"lisi": 26,
		"wangwu": 24,
	}

	// 遍历key和value
	for k,v := range scoreMap{
		fmt.Println(k,v)  // lisi 26 // wangwu 24 //zhangsan 24
	}

	// 只遍历key
	for k :=range scoreMap{
		fmt.Println(k, scoreMap[k])  // lisi 26 // wangwu 24 //zhangsan 24
	}

}
