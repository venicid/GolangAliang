package main

import "fmt"

func main()  {

	/**
	普通for
	 */
	var res int = 0
	for i:=1;i<=100;i++{
		res += i
	}
	fmt.Println("1-100的值：",res)

	/**
	外部定义 i
	 */
	var res2 int =0
	j := 1
	for j<= 100{
		res2 += j
		j++
	}
	fmt.Println(res2)

	/**
	模拟while
	 */
	res3 :=0
	m := 1
	for{   // 这里也等价 for ; ; {
		if m > 100{
			break   //break 就是跳出这个 for 循环
		}else{
			res3 += m
			m ++
		}
	}
	fmt.Println(res3)

	/**
	for range 键值循环

	0-a
	1-b
	2-c
	3-北
	6-上
	9-广

	*/
	var desc = "abc北上广"
	for index,val := range desc{
		fmt.Printf("%d-%c\n", index, val)
	}

	/**
	switch case
	 */
	score := "50"
	switch score {
	case  "90":
		fmt.Println("绩效是：A")
	case "70":
		fmt.Println("绩效是：B")
	case "60":
		fmt.Println("绩效是：C")
	case "50":
		fmt.Println("n+1")
	}


	/**
	break
	continue
	 */

	var data string = "abcdefg"
	var newString = ""
	for index,val := range data{
		if index == 3{
			continue
		}
		if val == 'e'{  // 字符
			break
		}
		newString  = newString + string(val)
	}
	fmt.Println(newString)  // abc
}
