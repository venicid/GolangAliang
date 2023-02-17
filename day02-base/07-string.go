package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	s1 := "hello"
	s2 := "world"

	s3 := `
		第1行
		第2行
		第3行
		`
	fmt.Println(s1 + s2)
	fmt.Println(s3)


	/*
	len
	 */
	var str = "hello world"
	fmt.Println(len(str)) // 11

	/*
	+ 拼接
	*/
	var str1 = "你好"
	var str2 = "golang"
	fmt.Println(str1 + str2)  // 你好golang

	/*
	strings.Split()
	*/
	var str3 = "123-456-789"
	var arr = strings.Split(str3, "-")
	fmt.Printf("arr: %v %T\n", arr, arr)  // arr: [123 456 789] []string

	/*
	strings.Join()
	*/
	var str4 = "110-119-120"
	var arr4 = strings.Split(str4, "-")
	var str5 = strings.Join(arr4, "*")
	fmt.Println(str5)    // 110*119*120


	/**
	// 字符 - 单引号
	// 字符串 - “双引号”
	*/
	a := 'a'
	// name := 'zhangsan'  // more than one character in rune literal
	name := "zhangsan"

	//当我们直接输出 byte（字符）的时候输出的是这个字符对应的码值
	fmt.Println(a)  // 97 这里输出的是 a 字符串的 ASCII值
	fmt.Printf("a: %T, %c %s\n", a, a, a)     // a: int32, a %!s(int32=97)
	fmt.Println(name)  // zhangsan

	//如果我们要输出这个字符，需要格式化输出
	fmt.Printf("的值是%c\n", a) // a的值是a


	/**
	byte和rune
	Go 语言的字符有以下两种
	uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符
	rune 类型，代表一个 UTF-8 字符

	字符串底层是一个byte数组，所以可以和[]byte类型相互转换。
	字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
	rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
	*/
	s := "美国第一"
	s_rune := []rune(s)

	fmt.Printf("s: %T, %c, %s\n", s, s, s)
	fmt.Printf("s_rune: %T, %c, %s\n", s_rune, s_rune, s_rune)

	fmt.Println("中国" + string(s_rune[2:]))
	/**
	s: string, %!c(string=美国第一), 美国第一
	s_rune: []int32, [美 国 第 一], [%!s(int32=32654) %!s(int32=22269) %!s(int32=31532) %!s(int32=19968)]
	中国第一
	 */


	/**
	遍历
	 */
	s = "hello world"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v: %c", s[i], s[i] )
	}
	// 104(h) 101(e) 108(l) 108(l) 111(o) 32( ) 229(å) 188(¼) 160() 228(ä) 184(¸) 137()

	fmt.Println() // 换行

	for _, i2 := range s {
		fmt.Printf("%v: %c",i2, i2 )
	}
	// 104(h) 101(e) 108(l) 108(l) 111(o) 32( ) 229(å) 188(¼) 160() 228(ä) 184(¸) 137()

	fmt.Println()


	/**
	修改字符串
	1. 转换成[]rune 或[]byte, 取值修改，转为string
	2. 无论哪种转换，都会重新分配内存，并复制字节数组
	 */
	oldS := "美国第一"
	oldRuneS := []rune(oldS)
	newS := "中国" + string(oldRuneS[2:])

	fmt.Printf("%v %c\n", oldRuneS, oldRuneS)
	// [32654 22269 31532 19968] [美 国 第 一]
	fmt.Println(newS)



	/**
	其他类型转string
	 */

	//1、int 转换成 string
	// int64 转 string
	var num1 int = 20
	sNum1 := strconv.Itoa(num1)
	fmt.Printf("sNum1:%v, %T\n", sNum1, sNum1)   // sNum1:20, string

	var num2 int64 = 8
	sNum2 := strconv.FormatInt(num2, 10)  /* 第二个参数10为 进制 */
	sNum21 := strconv.FormatInt(num2, 2)  /* 第二个参数 2进制 */
	fmt.Printf("sNum2:%v, %T\n", sNum2, sNum2)   // sNum2:8, string
	fmt.Printf("sNum21:%v, %T\n", sNum21, sNum21)  // sNum21:1000, string

	// 2、float 转 string
	/* 参数 1：要转换的值
	   参数 2：格式化类型
	   参数 3: 保留的小数点 -1（不对小数点格式化）
	   参数 4：格式化的类型
	*/
	var f float64 = 20.116623
	sFloat := strconv.FormatFloat(f, 'f', 3, 64)
	fmt.Printf("sFloat:%v, %T\n", sFloat, sFloat)  // sFloat:20.117, string

	// 3、bool 转 string
	var isError bool = false
	sIsError := strconv.FormatBool(isError)
	fmt.Printf("sIsError:%v, %T\n", sIsError, sIsError)   // sIsError:false, string



	/**
	string转其他类型
	 */
	var strNum1 string = "10"
	numStr1, _ := strconv.Atoi(strNum1)
	fmt.Printf("numStr1:%v， %T\n", numStr1, numStr1)

	var sFloat1 string = "1.111"
	floatS1,_ := strconv.ParseFloat(sFloat1, 64)
	fmt.Printf("floatS1:%v， %T\n", floatS1, floatS1)
}
