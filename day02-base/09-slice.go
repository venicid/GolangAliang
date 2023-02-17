package main

import "fmt"

/**
slice 可变长度 相同元素 序列
 */

func main(){

	/**
	定义
	 */
	var a []string

	fmt.Println(a) // []
	fmt.Println(a == nil)  // true
	fmt.Println(len(a))  // 0

	var b = []int{}
	fmt.Println(b)
	fmt.Println(b == nil)  //false
	fmt.Println(len(b))  // 0

	var c = []bool{true, false}  // [true false]
	fmt.Println(len(c))   // 2

	/**
	比较
	 */

	// 切片唯一合法的比较操作是和 nil 比较
	// 切片是引用类型，不支持直接比较，只能和 nil 比较
	// 一个 nil 值的切片并没有底层数组
	// 一个 nil 值的切片的长度和容量都是 0
	var d []int
	var e =  []string{}
	fmt.Println(d==nil)   // true
	fmt.Println(e==nil)   // false

	/**
	切片的本质: 长度和容量
	 */
	array := [8]int{0,1,2,3,4,5,6,7}
	fmt.Println(array)     // [0 1 2 3 4 5 6 7]

	s1 := array[0:5]
	fmt.Println(s1)   // [0 1 2 3 4]
	fmt.Println(len(s1))  // 5
	fmt.Println(cap(s1))   // 8

	s2 := array[3:6]
	fmt.Println(s2)   // [3 4 5]
	fmt.Println(len(s2))  // 3
	fmt.Println(cap(s2))   // 5


	/**
	遍历
	 */

	slice1 := []string{"a","b", "c", "d"}
	for i:=0; i<len(slice1);i ++{
		fmt.Println(i,slice1[i])
	}

	for key,value := range slice1{
		fmt.Println(key,value)
	}

	/**
	append 扩容
	 */
	numSlice:= []int{}
	for i:=0 ;i < 10;i++{
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	//[0] len:1 cap:1 ptr:0xc0000ac0d0
	//[0 1] len:2 cap:2 ptr:0xc0000ac0e0
	//[0 1 2] len:3 cap:4 ptr:0xc0000aa080
	//[0 1 2 3] len:4 cap:4 ptr:0xc0000aa080
	//[0 1 2 3 4] len:5 cap:8 ptr:0xc0000c6140
	//[0 1 2 3 4 5] len:6 cap:8 ptr:0xc0000c6140
	//[0 1 2 3 4 5 6] len:7 cap:8 ptr:0xc0000c6140
	//[0 1 2 3 4 5 6 7] len:8 cap:8 ptr:0xc0000c6140
	//[0 1 2 3 4 5 6 7 8] len:9 cap:16 ptr:0xc0000de080
	//[0 1 2 3 4 5 6 7 8 9] len:10 cap:16 ptr:0xc0000de080


	/**
	append 追加多个 追加切片 slice...
	切片合并： slice...
	 */
	var citySlice []string
	citySlice = append(citySlice, "beijing")
	citySlice = append(citySlice, "shanghai", "shenzhen")  // 追加多个元素

	// 追加切片  slice...
	secondSlice := []string{"hangzhou", "wuhan", "xian"}
	citySlice = append(secondSlice, secondSlice...)

	// 切片合并  slice...
	thirdSlice := []string{"tianjin", "nanchang", "xianyang"}
	citySlice = append(citySlice, thirdSlice...)

	fmt.Println(citySlice)


	/**
	查找
	 */
	nums := []int{ 22, 45,21,88, 99,10, 44}
	targetNum := 99
	flag := false
	for i:=0; i<len(nums); i++{
		if nums[i] == targetNum{
			flag = true
			break
		}
	}
	fmt.Println(flag)

	targetIndex := 3
	var targetValue int
	for i:=0; i<len(nums); i++{
		if  i == targetIndex{
			targetValue = nums[i]
			break
		}
	}
	fmt.Println(targetValue)


	/**
	删除
	*/
	// 要删除索引为 2 的元素
	nums1 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	nums1 = append(nums1[0:2], nums1[3:]...)
	fmt.Println(nums1)

	// 删除元素 99
	nums2 := []int{99, 31, 99, 22, 34, 35, 99, 37}
	targetNum2 := 99
	var newNums []int
	for i:=0 ; i< len(nums2); i++{
		if nums2[i] == targetNum2{
			continue
		}
		newNums = append(newNums, nums2[i])
	}
	fmt.Println(newNums)  // [31 22 34 35 37]


	/**
	修改
	*/
	nums3 := []int{99, 31, 99, 22, 34, 35, 99, 37}
	targetIndex3 := 3
	targetValue3 := 100
	oldValue3 := 99
	newValue3 := 88
	for i:=0;i<len(nums3); i++{
		if i ==targetIndex3 {
			nums3[i] = targetValue3
			continue
		}
		if nums3[i] == oldValue3{
			nums3[i] = newValue3
			continue
		}
	}
	fmt.Println(nums3)  // [88 31 88 100 34 35 88 37]

	/**
	copy
	 */
	nums4 := []int{99, 31, 99, 22, 34, 35, 99, 37}

	nums5 := nums4
	fmt.Println(nums4)
	fmt.Println(nums5)

	nums4[2] = 1
	fmt.Println(nums4)
	fmt.Println(nums5)

	//var nums6 []int  // 不生效
	nums6 := make([]int, len(nums4), (cap(nums4))*2)
	copy(nums6, nums4)
	nums4[0] = 1
	fmt.Println(nums4)
	fmt.Println(nums6)



}