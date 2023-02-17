package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	/**
	flag包实现了命令行参数的解析

	*/
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")
	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

}

/**
# 查看帮助
E:\helloGolang\src\day05_package> go run 04_flag.go --help
Usage of C:\Users\hua'wei\AppData\Local\Temp\go-build1618491053\b001\exe\04_flag.exe:
  -age int
        年龄 (default 18)
  -delay duration
        延迟的时间间隔 (default 10h0m0s)
  -married
        婚否 (default true)
  -name string
        姓名 (default "张三")

# 使用默认值
E:\helloGolang\src\day05_package>go run 04_flag.go
张三 18 false 0s
[]
0
0

# 非flag命令行参数
E:\helloGolang\src\day05_package>go run 04_flag.go a b c d
张三 18 false 0s
[a b c d]
4
0

# flag参数演示
E:\helloGolang\src\day05_package>go run 04_flag.go  alex 18 33 false 30 test
张三 18 true 10h0m0s
[alex 18 33 false 30 test]
6
0
*/
