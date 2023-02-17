package main

import (
	"fmt"
	"sync"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好golang")
		time.Sleep(time.Millisecond * 100)
	}
}

/**
WaitGroup
*/
var wg sync.WaitGroup // 第一步：定义一个计数器

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1() 你好golang")
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器-1 // 第三步：协程执行完毕，计数器-1
}

func test2() {
	for i := 0; i < 2; i++ {
		fmt.Println("test2() 你好golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器-1
}

func main() {

	/**
	开启goroutine，未等待子协程完成，直接退出
	*/
	go test() //表示开启一个协程

	for i := 0; i < 2; i++ {
		fmt.Println("main() 你好golang")
		time.Sleep(time.Millisecond * 100)
	}
	/**
	main() 你好golang
	test() 你好golang
	test() 你好golang
	main() 你好golang
	*/

	/**
	WaitGroup, 主进程通过该waitGroup等待协程完成
		sync.WaitGroup 内部维护一个计数器
		启动+1，调用Done()方法-1，Wait()等待并发任务完成
		计数器为0，表示并发任务全部完成
	*/
	fmt.Println("主线程开始...")
	wg.Add(1)  //协程计数器+1 第二步：开启一个协程计数器+1
	go test1() //表示开启一个协程
	wg.Add(1)  //协程计数器+1
	go test2() //表示开启一个协程
	wg.Wait()  //等待协程执行完毕... 第四步：计数器为0时推出

	fmt.Println("主线程退出...")
}
