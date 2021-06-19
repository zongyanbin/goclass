/*
	表单验证、自定义验证
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 结构体 ShouldBindJSON uri query?=
type PostParams struct {
	Name string `json:"name"`
	Age int `json:"age" binding:"required,mustBig"`
	Sex bool `json:"sex"`
}


//  mustBig 自定义验证
func mustBig(f1 validator.FieldLevel) bool{
	if(f1.Field().Interface().(int) <= 18){
		return  false
	}
	return true
}

func main(){
	r := gin.Default()

	// mustBig 表单验证挂载规则
	// 断言 binding.Validator.Engine().(*validator.Validate);
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// RegisterValidation 注册验证规则 mustBig
		v.RegisterValidation("mustBig", mustBig)
	}

	// POST ShouldBindQuery
	r.POST("/testBind", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBindJSON(&p)  // 把结构体用ShouldBindQuery绑定到p上，后面直接操作p
		if err!=nil{
			//fmt.Println(err.Error())  // 打印错误
			c.JSON(200,gin.H{
				"msg":"报错了",
				"data":gin.H{},
			})
		}else{
			c.JSON(200, gin.H{
				"msg":"成功了",
				"data":p,
			})
		}
	})

	// GET uri
	r.Run(":8081")
   }