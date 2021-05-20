/**
设置主键：gorm:"primary_key"
自定义字段名字：column:user_id
忽略： “—”
指定数据类型 type:varchar(100)
非空 notnull
创建索引：index
设置外键：Foreignkey
关联外键 AssociationForeignkey
多对多 many2many表名

 */

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
type User struct {
	gorm.Model
	Name string `gorm:"primary_key;column:user_name;type:varchar(100);"`
}

func (u User)TableName()string{
	return "qm_users"
}

func main(){
	dsn := "root:123456@tcp(127.0.0.1:3306)/gw?charset=utf8mb4&parseTime=True&loc=Local"
	open, err2 := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := open, err2
	if err !=nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&User{})
	//var user User
	fmt.Println(db)
	//db.First(&user,"name = ?","guomei")
}