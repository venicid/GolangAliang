package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	/**
	未加锁的问题
		最终的值不一致。存在一个goroutine取出num的值时做加法运算时，另一个goroutine也取出了num的值
	*/
	var num = 0
	for i := 0; i < 10000; i++ {
		go func(i int) {
			num++
			fmt.Printf("goroutine %d: num = %d\n", i, num)
		}(i)
	}
	time.Sleep(time.Second)  //主goroutine等待1秒以确保所有工作goroutine执行完毕
	fmt.Println("num:", num) // 最终的值不一致

	/**
	sync.Mutex是一个互斥锁，
		保证同时只有一个 goroutine 可以访问共享资源
	*/

	var lock sync.Mutex

	var num2 = 0
	for i := 0; i < 10000; i++ {
		go func(i int) {
			lock.Lock()
			num2++
			lock.Unlock()
			fmt.Printf("goroutine %d: num2 = %d\n", i, num2)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("num:", num) // 最终的值不一致
	fmt.Println("num2:", num2)

	/**
	waitGroup + Lock
	*/
	for i := 0; i < 10000; i++ {
		wg3.Add(1)
		go add()
	}
	wg3.Wait()
	fmt.Println("num3:", num3)
}

var num3 = 0
var lock3 sync.Mutex
var wg3 sync.WaitGroup

func add() {
	lock3.Lock()
	num3++
	lock3.Unlock()
	wg3.Done()
}
