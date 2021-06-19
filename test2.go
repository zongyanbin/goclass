/**
gorm

navicat 创建数据库
1.数据库连接
2.创建一个表用结构体 HelloWorld 三个字段
3.创建表、自动迁移、判断是否又这张表没有自动创建（AutoMigrate） gorm
4.创建表没有主键不规范，Gorm为我们提供规范
	id主键
	创建时间
	修改时间
	删除时间
	自动创建

type 在结构体里 struct {
	gorm.Model
}
新建表时需要把旧表删除，否则不会实现gorm

5、操作数据表
	增 :create （跟结构体指针地址）
	删除：delete(跟结构体指针地址)或者条件 会根据主键自动去查询单条或者根据条件删除多条
	该：update 更新单一数据还有updates更新数据中指定内容 save更新所有内容
    查：First （跟结构体实例指针地址） Find（跟结构体切片指针地址）
	条件：where or 填写简单的sql查询语句执行得到model
	模型：model  告诉你我要去那里查找、或者说我返回的结果是哪一个集，我在从集里做（增删改查）
gorm 如果在没有任何条件下传入一个空的结构体指针地址，那么他会以结构体指针对应的那张表的所有内容，作为一个model进行后续操作
*/


package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strings"
)

type HelloWorld struct {
	gorm.Model
	Name string
	Sex bool
	Age int
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gw?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil {
		log.Fatal(err.Error())
	}
	// 传递 helloworld结构体指针地址
	db.AutoMigrate(&HelloWorld{})

	// 创建
	//db.Create(&HelloWorld{
	//	Name:"rui",
	//	Sex:true,
	//	Age:80,
	//})

	// 查找 First 没有条件就查询表的所有数据的第一条 注意：可以指定条件
	/**
	var hello HelloWorld
	db.First(&hello,"name = ?","guomei")
	*/

	// 查找 Find 查找所有符合条件的内容 注意：接收切片
	var hello []HelloWorld
	//db.Find(&hello) //查询所有
	//db.Find(&hello, "age> ?", 200000) //查询带条件    age <21  很少这么写都写在where里
	//db.Where("age < ?", 21).Find(&hello) // where自动先创建一个model,然后在去find,会把model= Where("age < ?", 21)条件查出来   2、还有个Or()

	r:=gin.Default()
	r.GET("/tests", func(c *gin.Context) {
		//c.JSON(200,gin.H{
		//	"data":hello,
		//})
	})

	r.GET("/test", HanderDownload)

	r.PUT("", func(c *gin.Context) {
		_=c.BindJSON(&hello)
		c.JSON(200,gin.H{
			"data":hello,
		})
	})

	r.PUT("delete", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"delete_list":"hello",
		})
	})

	r.GET("update", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"delete_list":"delete",
		})
	})

	r.GET("")
	// 修改数据 (Update)
	//db.Where("id = ?",1).First(&HelloWorld{}).Update("name","qimiaoaaa") // 修改单挑数据

	//多条修改数据， (Updates)可以跟两个类型一个是结构体，一个是map对象
	//db.Where("id = ?",1).First(&HelloWorld{}).Updates(HelloWorld{
	//	Name: "qimiaoshuai",
	//	Age: 22,
	//}) // 修改多条数据

	//多条修改数据 用map对象  批量Find  Where("id in (?)",[]int{1,2}).First(&[]HelloWorld{})
	//db.Where("id in (?)",[]int{1,2}).Find(&[]HelloWorld{}).Updates(map[string]interface{}{
	//	"Name": "qimiaoshuai",
	//	"Sex":true,
	//	"Age": 22,
	//}) // 修改多条数据

	//删除 软删除 表里存在别的地方查不到
	//db.Where("id in (?)",[]int{1,2}).Delete(&HelloWorld{})
	//硬删除数据、Unscoped()
	//db.Where("id in (?)",[]int{1,2}).Unscoped().Delete(&HelloWorld{})
	r.Run(":8888")
}

func HanderDownload(c *gin.Context) {
	var hello []HelloWorld
	dsn := "root:123456@tcp(127.0.0.1:3306)/gw?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil {
		log.Fatal(err.Error())
	}
	// 传递 helloworld结构体指针地址
	db.AutoMigrate(&HelloWorld{})



	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A2", "我要下载一个excel文件")
	xlsx.SetCellValue("Sheet1", "A1", "有没有看到我帅气的脸庞")

	db.Find(&hello, "age> ?", 200000) //查询带条件    age <21  很少这么写都写在where里

	_=c.BindJSON(&hello)
	RefactorWrite(hello{})
	for k, v := range hello{
		fmt.Println(k,v.Name)
		axis := fmt.Sprintf("A%d", k+1)
		xlsx.SetCellValue("Sheet1", axis, v.Name)
		xlsx.SetCellValue("Sheet1", axis, v.Age)
	}



	//保存文件方式
 xlsx.SaveAs("./aaa.xlsx")

	//c.Header("Content-Type", "application/octet-stream")
	//c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	//c.Header("Content-Transfer-Encoding", "binary")
	//
	////回写到web 流媒体 形成下载
	//_ = xlsx.Write(c.Writer)

}

func RefactorWrite(records []*Record) {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")

	for i, t := range records {
		d := reflect.TypeOf(t).Elem()
		for j := 0; j < d.NumField(); j++ {
			// 设置表头
			if i == 0 {
				column := strings.Split(d.Field(j).Tag.Get("xlsx"), "-")[0]
				name := strings.Split(d.Field(j).Tag.Get("xlsx"), "-")[1]
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+1), name)
			}
			// 设置内容
			column := strings.Split(d.Field(j).Tag.Get("xlsx"), "-")[0]
			switch d.Field(j).Type.String() {
			case "string":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).String())
			case "int32":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).Int())
			case "int64":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).Int())
			case "bool":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).Bool())
			case "float32":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).Float())
			case "float64":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).Float())
			}
		}
	}

	xlsx.SetActiveSheet(index)
	// 保存到xlsx中
	err := xlsx.SaveAs("test_write.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

