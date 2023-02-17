package main

import "fmt"

func main() {

	/**
	select 类似于 switch
	*/
	// 在某些场景下我们需要同时从多个通道接收数据,这个时候就可以用到golang中给我们提供的select多路复用

	//1.定义一个管道 10个数据int
	intCh := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intCh <- i
	}

	//2.定义一个管道 5个数据string
	strCh := make(chan string, 20)
	for i := 0; i < 20; i++ {
		strCh <- "index=" + fmt.Sprintf("%d", i)
	}

	//使用select来获取channel里面的数据的时候不需要关闭channel
	for {
		select {
		case v := <-intCh:
			fmt.Println(v)
		case v2 := <-strCh:
			fmt.Println(v2)
		default:
			fmt.Println("数据获取完了")
			return //注意退出...
		}
	}

}
