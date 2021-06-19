/**
work
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

/**
四张主要表、一张关联表
	赛数据
 */
package main

// 引入gorm
import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 创建相关结构体
// 1对多 1个班级->多个学生
type Class struct {
	gorm.Model // 创建主键 创建时间 修改时间 删除时间
	ClassName string  //附加相应属性
	Students []Student
}
// 1对1 1个学生->1个身份证
type Student struct {
	gorm.Model // 创建主键 创建时间 修改时间 删除时间
	StudentName string //附加相应属性
	ClassID uint
	IDCard IDCard
	// 多对多
	Teachers []Teacher `gorm:"many2many:Student_Teachers"` // 一个学生有多个老师 主键自动关联
	//TeacherID uint //关联
}

// 身份证属于哪个学生
type IDCard struct {
	gorm.Model // 创建主键 创建时间 修改时间 删除时间
	StudentID uint
	Num int
}

type Teacher struct {
	gorm.Model // 创建主键 创建时间 修改时间 删除时间
	TeacherName string //附加相应属性
	// 老师学生 多对多 后看
	Students []Student `gorm:"many2many:student_teachers;"`
	//StudentID uint
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gw?charset=utf8mb4&parseTime=True&loc=Local"
	open, err2 := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := open, err2
	if err !=nil {
		log.Fatal(err.Error())
	}
	/** 手写代码关联
	db.AutoMigrate(&Class{}, &IDCard{}, &Student{}, &Teacher{})
	// 一个学生有身份证
	i:=IDCard{
		Num: 123456,
	}

	s:=Student{
		StudentName:"奇妙",
		IDCard:i,
		//Teachers: []Teacher{t}, // 关系表显示数据
	}

	t:=Teacher{
		TeacherName: "老师傅1",
		Students: []Student{s}, // 关系表显示数据
	}

	// 一个班级挂上多个学生
	c:=Class{
		ClassName: "奇妙的班级",
		Students: []Student{s},
	}
	_ = db.Create(&c).Error  //添加班级
	_ = db.Create(&t).Error	 //添加老师
	 */

	// 路由启动gin服务
	r:=gin.Default()
	// 路由获取POST请求接口、 func 接收上下文 以c来表示
	// 创建数据
	r.POST("/student", func(c *gin.Context) {
		var student Student
		_=c.BindJSON(&student)
		db.Create(&student)
		/**
		{
		    "ClassName":"qimiao",
		    "ClassID":1,
		    "IDCard":{
		        "Num":123456789
		    },
		    "TeacherNames":[{
		        "TeacherName": "老师1"
		    },{
		         "TeacherName": "老师2"
		    }]

		}
		 */
	})

	//查询数据 预加载 preload
	r.GET("/student/:ID", func(c *gin.Context) {
		id:=c.Param("ID") // 获取前端id
		var student Student
		_=c.BindJSON(&student)
		db.Preload("Teachers").Preload("IDCard").First(&student,"id=?",id)
		c.JSON(200,gin.H{
			"s":student,
		})
	})

	r.GET("/class/:ID", func(c *gin.Context) {
		id:=c.Param("ID") // 获取前端id
		var class Class
		_=c.BindJSON(&class)
		// students 优先做预加载 嵌套预加载
		db.Preload("Students").Preload("Students.Teachers").Preload("Students.IDCard").First(&class,"id=?",id)
		c.JSON(200,gin.H{
			"s":class,
		})
	})
 	r.Run(":111")

}
