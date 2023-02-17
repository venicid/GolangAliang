package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

/**
1、Many To Many
*/
// User 拥有并属于多种 language，`user_languages` 是连接表
type Group struct {
	gorm.Model
	Name    string
	Servers []Server `gorm:"many2many:group_servers;"`
}

type Server struct {
	gorm.Model
	Ip string
}

/**
2、反向引用
*/
//type Server struct {
//	gorm.Model
//	Ip     string
//	Groups []*Group `gorm:"many2many:group_servers;"`
//}

/**
3、重写外键
	对于 many2many 关系，连接表会同时拥有两个模型的外键
	可以仅使用其中的一个重写部分的外键、引用。
*/

// 连接表：user_languages
// foreign key: user_id, reference: users.id
// foreign key: language_id, reference: languages.id

type User struct {
	gorm.Model
	Profiles []Profile `gorm:"many2many:user_profiles;
						foreignKey:Refer;
						joinForeignKey:UserReferID;
						References:UserRefer;
						joinReferences;ProfileRefer"`
	Refer uint `gorm:"index:,unique"`
}
type Profile struct {
	gorm.Model
	Name      string
	UserRefer uint `gorm:"index,unique"`
}

/**
4、自定义第三张表
*/
type Person struct {
	ID      int
	Name    string
	Address []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
	ID   int
	Name string
}

type PersonAddress struct {
	PersonID  int `gorm:"primaryKey"`
	AddressID int `gorm:"primaryKey"`
	CreatedAt time.Time
	DeleteAt  gorm.DeletedAt
}

func (PersonAddress) BeforeCreate(db *gorm.DB) (err error) {
	// 修改 Person 的 Addresses 字段的连接表为 PersonAddress
	// PersonAddress 必须定义好所需的外键，否则会报错
	err = db.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{})
	if err != nil {
		fmt.Println("err", err)
	}
	return nil
}

func main() {
	// 1、连接数据库
	dsn := "root:root112358@tcp(150.158.171.205:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 2、自动创建多对的表结构
	//db.AutoMigrate(Group{}, Server{})

	db.AutoMigrate(Person{}, Address{})

	// 2、添加数据
	//person := Person{
	//	ID:   2,
	//	Name: "lisi",
	//	Address: []Address{
	//		{ID: 1, Name: "beijing"},
	//		{ID: 3, Name: "shenzhen"},
	//	},
	//}
	//db.Create(&person)

	/**
	多对多Preload
	*/
	// 查找全量list
	persons := []Person{}
	db.Preload("Address").Find(&persons)

	personsByte, _ := json.Marshal(&persons)
	fmt.Println(string(personsByte))

	// 查找单个用户map
	person := Person{Name: "zhangsan"}
	db.Preload("Address").Find(&person)

	personByte, _ := json.Marshal(&person)
	fmt.Println(string(personByte))
	// {"ID":1,"Name":"zhangsan","Address":[{"ID":1,"Name":"beijing"},{"ID":2,"Name":"shanghai"}]}

}
