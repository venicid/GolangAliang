package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Admin 表的结构体ORM映射
type Admin struct {
	Id       int64 `json:"id" gorm:"primary_key"`
	Username string
	Password string
}

func main() {

	// 1、连接数据库
	dsn := "root:root112358@tcp(150.158.171.205:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db) // &{0xc00020c510 <nil> 0 0xc000216380 1}

	// 2、自动创建表
	db.AutoMigrate(Admin{})

	/**
	增删改查
	*/
	// 1、增
	db.Create(&Admin{
		//Id: 3,
		Username: "zhangsan",
		Password: "123456",
	})

	// 2、改
	db.Model(Admin{
		Id: 3,
	}).Update("username", "lisi")

	// 3、查
	// 3.1 过滤查询
	u := Admin{Id: 3}
	db.First(&u)
	fmt.Println(u)
	// 3.2 查询所有数据
	users := []Admin{}
	db.Find(&users)
	fmt.Println(users) // [{2 zhangsan 123456} {3 lisi 123456}]

	// 4、删
	// 4.1 删除 id = 3 的用户
	db.Delete(&Admin{Id: 3})
	// 4.2 条件删除
	db.Where("username = ?", "zhangsan").Delete(&Admin{})

}
