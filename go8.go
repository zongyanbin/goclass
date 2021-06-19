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

/**
学生 老师 班级 身份证号码
学生->身份证 1对1
学生->老师  多对多
班级->学生 ·1对多
*/

package main
// 引入gorm
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
// 创建User表自动生成两个主键
type User struct {
	gorm.Model   // 设置主键
	Name string `gorm:"primary_key;column:user_name;type:varchar(100)"` // 名字也设置主键  自定义字段名 column:user_name
}

// gorm生成是Users
// 自定义表名 qm_users user挂载一个方法  支持条件语法
func (u User) TableName() string{
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