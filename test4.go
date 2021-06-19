package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type AdminUser struct {
	gorm.Model
	Username string
	Age uint8
	Order []Order
}

type Order struct {
	gorm.Model
	AdminUser uint
	Price float64
	AdminUserID uint
	OrderType []OrderType
}

type OrderType struct {
	gorm.Model
	OrderID uint
	OrderName string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gw?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	// 传递 helloworld结构体指针地址
	//db.AutoMigrate(&Order{},OrderType{},AdminUser{})
	var _ = []AdminUser{
		{Username: "jinzhu1"},
		{Username: "jinzhu2"},
		{Username: "jinzhu3"},
		{Username: "jinzhu4"},
		{Username: "jinzhu5"},
	}
	//db.Select("Username").Create(&adminuser)
	//	db.Create(&adminuser)
	//	db.Debug().Preload("Order").Find(&adminuser)

	// map 二位数组
	//db.Model(&AdminUser{}).Create([]map[string]interface{}{
	//	{"Username": "jinzhu4"},
	//	{"Username": "jinzhu5"},
	//	{"Username": "jinzhu6"},
	//})

	// map 一位数组
	db.Model(&AdminUser{}).Create(map[string]interface{}{
		"Username": "我是一位数组map",
	})
	db.Model(&AdminUser{}).Create(map[string]interface{}{
		"Username": "map",
	})
	var adminuser = AdminUser{}
	fmt.Println("---")
	result := db.First(&adminuser)
	//result.RowsAffected // 返回找到的记录数
	fmt.Println(result.RowsAffected)
}

