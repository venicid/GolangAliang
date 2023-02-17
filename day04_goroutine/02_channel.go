package main

import "fmt"

func main() {

	/**
	channel，引用类型，管道
	*/
	//channel类型, var 变量 chan 元素类型
	var ch1 chan int           // 声明一个传递整型的管道
	var ch2 chan bool          // 声明一个传递布尔型的管道
	var ch3 chan []int         // 声明一个传递 int 切片的管道
	fmt.Println(ch1, ch2, ch3) // <nil> <nil> <nil>

	// 创建channel, 格式如下： make(chan 元素类型, 容量)
	ch11 := make(chan int, 10)    // 创建一个能存储 10 个 int 类型数据的管道
	ch22 := make(chan bool, 4)    // 创建一个能存储 4 个 bool 类型数据的管道
	ch33 := make(chan []int, 3)   // 创建一个能存储 3 个[]int 切片类型数据的管道
	fmt.Println(ch11, ch22, ch33) // 0xc000110000 0xc0000180e0 0xc00004c060

	// 操作
	// 1、创建channel
	ch := make(chan int, 5)

	// 2、向channel放入数据
	ch <- 10
	ch <- 12
	fmt.Println("发送成功", ch)

	// 3、向channel取值
	v1 := <-ch
	fmt.Println(v1)
	v2 := <-ch
	fmt.Println(v2)

	// 4、空channel取值报错
	//v3 := <-ch
	//fmt.Println("v3", v3)

	/**
	优雅的从channel取值
	*/
	c := make(chan string)

	go closeDemo(c)

	for i := range c { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}

}

func closeDemo(c chan string) {
	c <- "a"
	c <- "b"
	close(c) // 过close函数关闭通道来告知从该通道接收值的goroutine 停止等待
}
