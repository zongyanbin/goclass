package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type HelloWorld struct {
	// 规范  自动添加主键 新增时间  更新时间 删除时间
	gorm.Model
	Name string
	Age int
	Sex bool
}

func main(){
	db := Mysql()
	db.AutoMigrate(&HelloWorld{})
	
	//// 增加
	//db.Create(&HelloWorld{
	//	Name: "ran",
	//	Sex:false,
	//	Age: 50,
	//})

	//var hello HelloWorld
	// 查询
	//db.First(&hello,"name=?","grant")
	//db.Where("age < ?",20).First(&hello)
	//fmt.Println(hello)

	// 修改 修改多条update(结构体， 或者对象)  不会修改空属性
	//db.Where("id = ?",1).First(&HelloWorld{}).Updates(HelloWorld{
	//	Name: "qimiaoshuai",
	//	Age: 26,
	//})

	db.Where("id = ?",1).First(&HelloWorld{}).Updates(map[string]interface{}{
		"Name": "qimiaoshuai",
		"Sex":false,
		"Age": 26,
	})

}

// 封装mysql
func Mysql() *gorm.DB{
	dsn := "root:123456@tcp(127.0.0.1:3306)/gw?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}