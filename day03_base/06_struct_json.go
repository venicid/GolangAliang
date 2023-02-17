package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// 构造函数
type Lvs struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Env  string `json:"env"`
	Port int    `json:"port"`
}

func newLvs(lvs Lvs) Lvs {
	lvs.Env = strings.ToUpper(lvs.Env)
	return lvs
}

// 组合与嵌入
type App struct {
	Id    int            `json:"id"`
	Name  string         `json:"name"`
	Other Other          // 组合
	Extra `json:"extra"` // 嵌入（推荐）

	Groups []Group `json:"groups"`
}

//func (a *App) getInfo() string {
//	return a.Name
//}

type Other struct {
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	CreateUser string `json:"create_user"`
	UpdateUser string `json:"update_user"`
}

func (o *Other) getInfo() string {
	return o.CreateTime
}

type Extra struct {
	Level int    `json:"level"`
	IsOut bool   `json:"is_out"`
	Desc  string `json:"desc"`
}

func (e *Extra) getInfo() string {
	return e.Desc
}

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (g *Group) getInfo() string {
	return g.Name
}

func main() {
	/**
	1、struct组成的切片 ---> json字符串
	*/
	var lvs1 = Lvs{
		Id:   1,
		Name: "fat-test-inner.com",
		Env:  "fat",
		Port: 80,
	}
	var lvs2 = Lvs{
		Id:   2,
		Name: "uat-test-inner.com",
		Env:  "uat",
		Port: 80,
	}

	var lvsList = []Lvs{lvs1, lvs2}
	fmt.Println(lvsList) // [{1 fat-test-inner.com fat 80} {2 uat-test-inner.com uat 80}]

	lvsListByte, _ := json.Marshal(lvsList)
	fmt.Printf(string(lvsListByte))
	// [{"id":1,"name":"fat-test-inner.com","env":"fat","port":80},{"id":2,"name":"uat-test-inner.com","env":"uat","port":80}]{3 pre-test-inner.com PRE 80}

	/**
	2、构造函数
	*/
	var lvs3 = newLvs(
		Lvs{
			Id:   3,
			Name: "pre-test-inner.com",
			Env:  "pre",
			Port: 80,
		})
	fmt.Println(lvs3) // {3 pre-test-inner.com PRE 80}

	/**
	3、struct高级: 组合、转发
		组合，通过结构实现组合
		转发，通过名为嵌入的特殊语言特性，实现方法转发（推荐）
	*/
	var group1 = Group{
		Id:   1,
		Name: "AppCMDBApi-SH01",
	}
	var group2 = Group{
		Id:   2,
		Name: "AppCMDBApi-AL01",
	}

	var app1 = App{
		Id:   1,
		Name: "AppCMDBApi",
		Extra: Extra{
			Level: 3,
			IsOut: false,
			Desc:  "cmdb后端",
		},
		Groups: []Group{group1, group2},
	}
	fmt.Printf("%#v\n", app1)

	app1Byte, _ := json.Marshal(app1)
	app1JsonStr := string(app1Byte)
	fmt.Printf("app1: %v\n", app1JsonStr)
	/**
	app1:
	{"id":1,"name":"AppCMDBApi",
	"extra":{"level":3,"is_out":false,"desc":"cmdb后端"},
	"groups":[{"id":1,"name":"AppCMDBApi-SH01"},{"id":2,"name":"AppCMDBApi-AL01"}]
	}
	*/

	fmt.Printf("%v的应用等级是%v\n", app1.Name, "S"+strconv.Itoa(app1.Extra.Level)) // AppCMDBApi的应用等级是S3
	fmt.Printf("%v的group数量是%v\n", app1.Name, len(app1.Groups))                // AppCMDBApi的group数量是2

	fmt.Println(app1.Other.getInfo())
	fmt.Println(app1.getInfo()) // 有歧义的方法选择器（主结构体没有该方法，自动调用子结构体的方法）
	fmt.Println(app1.Extra.getInfo())
}
