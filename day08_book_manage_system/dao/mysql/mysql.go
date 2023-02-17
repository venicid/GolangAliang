package mysql

import (
	"day08_book_manage_system/model"
	"fmt"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	// 1、连接数据库
	dsn := "root:root112358@tcp(150.158.171.205:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("初始化mysql连接错误, err:", err)
	}
	DB = db

	if err := DB.AutoMigrate(model.User{}, model.Book{}); err != nil {
		fmt.Println("自动创建表结构失败：", err)
	}
}
