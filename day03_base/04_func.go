package main

import "fmt"

func sum(x, y int) (z int){
	return x+y
}


// 全局变量
//username1 := "admin"  // 未声明
var username = "admin"

// 局部变量
func test(){
	var testUser string = "test"
	fmt.Println(testUser)
}


func main(){
	/**
	函数：定义+返回值
	 */
	var m = 4
	var n = 5
	r := sum(m,n)
	fmt.Printf("%v+%v=%v",m,n,r)


	/**
	变量：全局，局部，for中变量
	 */
	fmt.Println("username:", username)


	test()
	// 这里name是函数test的局部变量，在其他函数无法访问
	//fmt.Println(testUser)


	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量 i 只在当前 for 语句块中生效
	}
	// fmt.Println(i) //此处无法使用变量 i


}

