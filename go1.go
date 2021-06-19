package main

import "github.com/gin-gonic/gin" // 1.引入gin

func main(){
	r := gin.Default()   // 2.gin 路由启动jin服务  携带基础中间件启动
	r.GET("/ping", func(c *gin.Context){ // 3.路由接收GET请求/接口 func里接收gin的上下文 以c来表示
		c.JSON(200, gin.H{ // 4. c.JSON把这条信息返回去
			"sucess":true,
		})
	})
	r.Run(":1010")
}

func add(){
    r :=gini.Default();
}
