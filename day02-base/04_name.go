package main

import "fmt"

/**
命名规范

Go是一门区分大小写的语言。

以大写字母开头
	可以被外部包的代码所使用
以小写字母开头（像面向对象语言中的 private ）
	对包外是不可见的，但是他们在整个包的内部是可见并且可用的

包名称
	为小写单词，不要使用下划线或者混合大小写
	package domain
	package main

文件命名
	为小写单词，使用下划线分隔
	approve_service.go

结构体命名
	采用驼峰命名法，
	首字母根据访问控制大写或者小写struct 申明和初始化格式采用多行

	type MainConfig struct {
		Port string `json:"port"`
		Address string `json:"address"`
	}

接口命名
	与结构体命名类似，以 “er” 作为后缀，例如 Reader , Writer 。
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

变量命名
	一般遵循驼峰法，首字母根据访问控制原则大写或者小写
	特有名词时，需要遵循以下规则：如果变量为私有，且特有名词为首个单词，则使用小写
	bool 类型 ，则名称应以 Has, Is, Can 或 Allow 开头

	var isExist bool
	var hasConflict bool
	var canManage bool
	var allowGitHook bool

常量命名
	全部大写字母组成，并使用下划线分词
	如果是枚举类型的常量，需要先创建相应类型：

	const APP_URL = "https://www.baidu.com"

	type Scheme string const (
		HTTP Scheme = "http"
		HTTPS Scheme = "https"
	)
*/

func main()  {
	fmt.Println(1111)
}