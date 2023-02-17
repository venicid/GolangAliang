package main

import (
	"day08_book_manage_system/dao/mysql"
	"day08_book_manage_system/router"
	"fmt"
)

func main() {
	// 初始化mysql连接
	mysql.InitMysql()

	// 初始化，路由分层
	r := router.InitRouters()

	err := r.Run(":8888")
	if err != nil {
		fmt.Println("gin start failed, err:", err)
	}
	fmt.Println("gin start success!!!")
}
