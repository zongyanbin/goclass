/**
get post delete put请求以及获取参数的方式
不一样的四个接口
GET 不可以从body JSON拿数据
POST PUT DELETE 可以从JSON拿数据
 */
package main

import "github.com/gin-gonic/gin"

func main()  {
	// 声明一次就可以通用
	r := gin.Default()  // 路由启动gin服务  携带基础中间件启动

	// get 方式 ：获取uri里参数
	r.GET("/path/:id", func(c *gin.Context) {  // 路由接收GET请求接口 func 接收gin上下文 以c来表示
		id := c.Param("id") // 获取uri里的参数
		user := c.DefaultQuery("user","王二小") // 获取Query传参 分配默认值
		pwd := c.Query("pwd") // 获取Query传参
		c.JSON(200, gin.H{ // c.JSON返回数据
			"id":id,
			"user":user,
			"pwd":pwd,
		})
	})

	//POST方式 :获取参数
	r.POST("/path", func(c *gin.Context) { // 路由接收POST请求接口 func 接收gin上下文 以c来表示
		user :=c.DefaultPostForm("user","默认值post")
		pwd :=c.PostForm("pwd")
		c.JSON(200, gin.H{ // c.JSON返回数据
			"user":user,
			"pwd":pwd,
		})
	})

	// DELETE 方式 uri用到
	r.DELETE("/path/:id", func(c *gin.Context) { //路由接收 DELETE 请求接口 func 接收gin上下文 以c来表示
		id := c.Param("id")
		c.JSON(200,gin.H{ // c.JSON返回数据
			"id":id,
		})
	})

	// put 修改方式
	r.PUT("/path", func(c *gin.Context) { //路由接收 put 请求接口接口 func 接收gin上下文 以c来表示
		user :=c.DefaultPostForm("user","默认值post")
		pwd :=c.PostForm("pwd")
		c.JSON(200, gin.H{ // c.JSON返回数据
			"user":user,
			"pwd":pwd,
		})
	})
	r.Run(":1012")
}
