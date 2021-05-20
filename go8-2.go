/**
学生 老师 班级 身份证号码
学生->身份证 1对1
学生->老师  多对多
班级->学生 ·1对多
 */

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
type  Class struct{
	gorm.Model
	ClassName string
	//StudentName string
	//班级里有多个学生
	Students []Student
}

type Student struct {
	gorm.Model
	StudentName string
	//学生属于哪个班级 ClassID
	ClassID uint
	//一个学生都要有一个ID号
	IDCard IDCard
	//一个学生有多个老师
	Teachers []Teacher `gorm:"many2many:student_teachers;"`
	//学生想知道自己有几个老师
	//TeacherID uint
}

type  IDCard struct {
	gorm.Model
	StudentID uint
	Num int
}

type Teacher struct {
	gorm.Model
	TeacherName string
	//老帅有多个学生
//	StudentID uint
	Students []Student  `gorm:"many2many:student_teachers;"`
}

func main()  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gw?charset=utf8mb4&parseTime=True&loc=Local"
	open, err2 := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := open, err2
	if err !=nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&Teacher{}, &Student{},&Class{},&IDCard{})
/**
	i:=IDCard{
		Num:123456,
	}
	//赛数据 1对1
	s:=Student{
		StudentName: "qm",
		IDCard:i,
	}

	t:=Teacher{
		TeacherName: "老师傅",
	}
	c:=Class{
		ClassName: "奇妙的班级",
		//班级里有学生, 一个班级挂上多个学生
		Students: []Student{s},
	}

	//_ = db.Create(&c).Error
	//_ = db.Create(&t).Error
**/
	// 用gin创建服务
	r := gin.Default()
	r.POST("/student", func(c *gin.Context) {
		var student Student
		_ =c.BindJSON(&student) // student已经拿到所有相关属性
		db.Create(&student)
	})
	r.Run(":8888")
}


//
//{
//"StudentName":"qm",
//"ClassID":1,
//"IDCard":{
//"Num":5555
//},
//Teachers:[{
//"TeacherName":"老帅1"
//},{
//"TeacherName":"老帅2"
//}]
//
//}