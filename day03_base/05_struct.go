package main

import (
	"encoding/json"
	"fmt"
)

type server struct {
	id  int
	ip  string
	env string
}

func (s server) getEnv() string {
	return s.env
}

func (s server) setEnv(env string) {
	s.env = env
}

func (s *server) setEnv2(env string) {
	s.env = env
}

type Slb struct {
	ID         int
	Name       string
	Env        string
	createTime string // 私有属性不能被json包访问，序列化
}

type Nginx struct {
	ID         int    `json:"id"` //通过指定 tag 实现 json 序列化该字段时的 key
	Name       string `json:"name"`
	Env        string
	createTime string
}

func main() {

	/**
	1、三种实例化方式
	*/
	/**
	基本实例化
	*/
	var s1 server
	s1.id = 1
	s1.ip = "192.168.12.12"
	s1.env = "fat"

	fmt.Printf("s1=%v\n", s1)  // s1={1 192.168.12.12 env}
	fmt.Printf("s1=%#v\n", s1) // s1=main.server{id:1, ip:"192.168.12.12", env:"env"}
	/**
	1. %v    只输出所有的值
	2. %+v 先输出字段名字，再输出该字段的值
	3. %#v 先输出结构体名字值，再输出结构体（字段名字+字段的值）
	*/

	/**
	new 实例化
		- 返回 结构体的地址
		- s2是 结构体指针  s2=&main.server{}
	*/
	var s2 = new(server)
	s2.id = 2
	s2.ip = "8.8.8.8"
	s2.env = "pro"
	fmt.Printf("s2=%#v\n", s2) // s2=&main.server{id:2, ip:"8.8.8.8", env:"pro"}

	// 对比
	fmt.Printf("s1.ip=%v\n", s1.ip)
	fmt.Printf("s2.ip=%v\n", s2.ip)
	//fmt.Printf("s1.ip=%v\n", (*s1).ip)  // 不支持
	fmt.Printf("s2.ip=%v\n", (*s2).ip) // 对结构体指针直接使用.来访问结构体的成员

	/**
	键值对 实例化
	*/
	var s3 = server{
		id:  3,
		ip:  "192.168.23.22",
		env: "uat", // 注意：最后一个属性的,要加上
	}
	fmt.Printf("s3=%#v\n", s3) // s3=main.server{id:3, ip:"192.168.23.22", env:"uat"}

	/**
	2、方法, 接收者
	*/
	/**
	值类型
	*/
	var s4 = server{
		id:  3,
		ip:  "1.1.1.1",
		env: "fat",
	}

	s4.setEnv("FAT")
	fmt.Println(s4.getEnv()) // fat , 无法修改接收者变量本身。

	/**
	指针类型
	*/
	s4.setEnv2("UAT")
	fmt.Println(s4.getEnv()) // UAT ,修改都是有效的。 类似于 this 或者 self

	/**
	3、序列化与反序列化
		jsonString struct
		json.Marshal()  json.UnMarshal()
	*/

	/**
	struct  -->  jsonString
	*/
	// 私有属性不能被 json 包访问
	// 小写的，代表私有
	var s5 = server{
		id:  5,
		ip:  "1.1.1.1",
		env: "fat",
	}
	fmt.Printf("s5 %#v\n", s5)

	var s5Byte, _ = json.Marshal(s5)
	fmt.Printf("s5Byte 类型: %T, 值: %v\n", s5Byte, s5Byte)

	s5JsonStr := string(s5Byte)
	fmt.Printf("s5JsonStr 类型: %T, 值: %v\n", s5JsonStr, s5JsonStr)
	/**
	s5 main.server{id:5, ip:"1.1.1.1", env:"fat"}
	s5Byte 类型: []uint8, 值: [123 125]
	s5JsonStr 类型: string, 值: {}
	*/

	// 公有属性，public的struct
	var slb1 = Slb{
		ID:         1,
		Name:       "slb-test-inner.com",
		Env:        "fat",
		createTime: "2022-10-29 10:27:50",
	}
	fmt.Printf("slb1: %#v\n", slb1)

	var slb1Byte, _ = json.Marshal(slb1)
	fmt.Printf("slb1Byte 类型: %T  值: %v\n", slb1Byte, slb1Byte)

	var slb1JsonStr = string(slb1Byte)
	fmt.Printf("slb1JsonStr 类型: %T  值: %v\n", slb1JsonStr, slb1JsonStr)
	/**
	slb1: main.Slb{ID:1, Name:"slb-test-inner.com", Env:"fat", createTime:"2022-10-29 10:27:50"}
	slb1Byte 类型: []uint8  值: [123 34 73 68 34 58 49 44 34 78 97 109 101 34 58 34 115 108 98 45 116 101 115 116 45 105 110 110 101 114 46 99 111 109 34 44 34 69 110 118 34 58 34 102 97 116 34 125]
	slb1JsonStr 类型: string  值: {"ID":1,"Name":"slb-test-inner.com","Env":"fat"}
	*/

	/**
	jsonString  --> struct
	*/
	// go不支持这种类型, dict, json
	//var json1 = {"ID":1,"Name":"slb-test-inner.com","Env":"fat"}
	// 非法json字符串
	//var jsonString1 string = `{"ID":1"","Name":"slb-test-inner.com","Env":"fat"}`

	var jsonString1 string = `{"ID":1,"Name":"slb-test-inner.com","Env":"fat"}`
	var slb2 Slb

	err := json.Unmarshal([]byte(jsonString1), &slb2) //  &slb2 取地址，修改是有效的
	if err != nil {
		fmt.Printf("unmarshal err: %v\n", err)
	}

	fmt.Printf("slb2: %#v\n", slb2)
	fmt.Println("slb2.name:", slb2.Name)
	// slb2: main.Slb{ID:1, Name:"slb-test-inner.com", Env:"fat", createTime:""}
	// slb2.name: slb-test-inner.com

	/**
	struct tag
		结构体的元信息，运行时，通过反射机制读取
		key1:"value1" key2:"value2
	*/
	// tag结构体 --> jsonString

	var nginx1 = Nginx{
		ID:         1,
		Name:       "fat-test-inner.com",
		Env:        "fat",
		createTime: "2022-10-29 10:47:53",
	}
	fmt.Printf("nginx1:%#v\n", nginx1)

	nginx1Byte, _ := json.Marshal(nginx1)
	nginx1String := string(nginx1Byte)

	fmt.Printf("nginx1Byte 类型: %T  值: %v\n", nginx1Byte, nginx1Byte)
	fmt.Printf("nginx1String 类型: %T  值: %v\n", nginx1String, nginx1String)
	// nginx1Byte 类型: []uint8  值: [123 34 105 100 34 58 49 44 34 110 97 109 101 34 58 34 102 97 116 45 116 101 115 116 45 105 110 110 101 114 46 99 111 109 34 44 34 69 110 118 34 58 34 102 97 116 34 125]
	// nginx1String 类型: string  值: {"id":1,"name":"fat-test-inner.com","Env":"fat"}

	// jsonString -- > tag结构体
	nginx2String := ` {"id":1,"name":"fat-test-inner.com","Env":"fat"}`
	var nginx2 Nginx
	err = json.Unmarshal([]byte(nginx2String), &nginx2)
	if err != nil {
		fmt.Printf("unmarshal err: %v\n", err)
	}

	fmt.Printf("nginx2: %#v\n", nginx2)
	fmt.Println("nginx2.name:", nginx2.Name)
	// nginx2: main.Nginx{ID:1, Name:"fat-test-inner.com", Env:"fat", createTime:""}
	// nginx2.name: fat-test-inner.com
}
