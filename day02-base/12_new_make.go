package main

import "fmt"

func main()  {
	/**
	场景：执行错误
		引用类型，需要分配内存空间
	panic: assignment to entry in nil map

	goroutine 1 [running]:
	 */
	//var userInfo map[string]string
	//userInfo["username"] = "alex"
	//fmt.Println(userInfo)


	/**
	make与new的比较
		相同：创建并分配类型的内存
		区别：
			make：创建slice，map、channel等数据结构
			new: 为类型申请一片内存空间，并返回这片内存的指针
	 */

	a := make([]int, 3, 10)  // // 切片长度为 3，预留空间长度为 10
	a = append(a,1)
	a = append(a,2)
	fmt.Printf("%v-%T\n", a,a)   // [0 0 0 1 2]-[]int

	var b = new([]int)
	//b = b.append(b,2) // 返回的是内存指针，所以不能直接 append
	*b = append(*b ,3 )  // 必须通过 * 指针取值，才能进行 append 添加

	fmt.Printf("%v-%T\n", b,b)  // &[3]-*[]int  内存的指针---内存指针


	/**
	make 函数
		只用于 chan、map 以及 slice 的内存创建，返回的三个类型本身
	*/
	s1 := make([]int, 3, 10)   // 为slice分配内存的时候，应当尽量预估到slice可能的最大长度
	m1 := make(map[string]string)
	c1 := make(chan int,1 )
	fmt.Println(s1,m1,c1)  // [0 0 0] map[] 0xc0000180e0



	/**
	new 函数
		分配内存空间，返回指针
	 */

	// 系统默认的数据类型，分配空间

	// 1. 实例化int
	age := new(int)
	*age = 11
	fmt.Println(*age)  // 11

	// 2、实例化slice
	li := new([]int)
	*li = append(*li, 1)
	fmt.Println(*li)  // [1]

	// 3. 实例化map
	m := new(map[string]string)
	fmt.Println(*m)  // map[]

	*m = map[string]string{}
	(*m)["m"] = "alex"
	(*m)["username"] = "jack"
	fmt.Println(*m)   // map[m:alex username:jack]


	// 自定义类型使用new

	var s *Student
	s = new(Student)   //分配空间
	s.name = "pig"
	fmt.Println(s)   // &{pig 0}

}

type Student struct {
	name string
	age int
}