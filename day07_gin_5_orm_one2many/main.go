package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
constraint:OnUpdate:CASCADE 【当User表更新，也会同步给CreditCards】 // 外键约束
OnDelete:SET NULL 【当User中数据被删除时，CreditCard关联设置为 NULL，不删除记录】
*/

type User struct {
	gorm.Model               // 添加基础信息，创建时间，删除时间，更新时间，id
	Username    string       `json:"username" gorm:"column:username"`
	CreditCards []CreditCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	// 1、连接数据库
	dsn := "root:root112358@tcp(150.158.171.205:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 创建表结构
	//db.AutoMigrate(User{}, CreditCard{})

	/**
	1、创建一对多
	*/
	//user := User{
	//	Username: "zhangsan",
	//	CreditCards: []CreditCard{
	//		{Number: "0001"},
	//		{Number: "0002"},
	//	},
	//}
	//db.Create(&user)

	// 2、为已存在用户添加信用卡
	//u := User{Username: "zhangsan"}
	//db.First(&u)
	////fmt.Println(u.Username)
	//db.Model(&u).Association("CreditCards").Append([]CreditCard{
	//	{Number: "0003"},
	//})

	/**
	查找关联：一对多Association
		把 User 查询好, 然后根据 User 定义中指定的
		AssociationForeignKey 去查找 CreditCard
	*/
	// 3.查找 用户名为 zhangsan 的所有信用卡信息
	u1 := User{
		Username: "alex",
	}
	// Association必须要先查出User才能关联查询对应的 CreditCard
	db.First(&u1)
	err := db.Model(&u1).Association("CreditCards").Find(&u1.CreditCards)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", u1)
	userByte, _ := json.Marshal(&u1)
	fmt.Printf("%#v\n", userByte)
	fmt.Println(string(userByte))

	/**
	预加载：一对多Preload
		使用 Preload 方法, 在查询 User 时先去获取 CreditCard 的记录

	分为：单个查询or批量查询
	*/
	users := User{Username: "alex"} // 只查找张三用户的信用卡信息  {}  ，存在多个张三，只返回第一个数据，不准确
	//users := []User{} // 查找全量 [{},{}]

	db.Preload("CreditCards").Find(&users)

	userByte, _ = json.Marshal(&users)
	fmt.Println(string(userByte))

}
