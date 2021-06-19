/**
把所有路由块加上前缀，证明在同一个分组下面
同一个组下面可以创建多个方法

分组：
1、代码清晰，和代码相关的代码，
2、更下方便的管理路由添加中间件
脑图：

d
https://naotu.baidu.com/file/e16040cf726f1e20041b5f0957752ad3?token=a2361369a32ac4f3
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 不接受入参只接受回参
func middel()gin.HandlerFunc{
	return func(c *gin.Context) {
		fmt.Println("我在方法前，我是1")
		c.Next()
		fmt.Println("我在方法后，我是1")
	}
}

func middeltwo()gin.HandlerFunc{
	return func(c *gin.Context) {
		fmt.Println("我在方法前，我是2")
		c.Next()
		fmt.Println("我在方法后，我是2")
	}
}

func main()  {
	r := gin.Default()  // 路由gin服务启动

	// 对一个分组用中间件 两种方式  Use(middel(),middeltwo())
	v1 := r.Group("v1").Use(middel()).Use(middeltwo())

	v1.GET("test", func(c *gin.Context) {
		fmt.Println("我在方法内部")
		c.JSON(200,gin.H{
			"success":true,
		})
	})

	r.Run(":8080")
}