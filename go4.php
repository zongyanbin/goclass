/**
 	单文件上传
	文件的上传与返回
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main(){
	r := gin.Default() // 路由启动jin服务、 携带基础中间件启动
	// 给表单限制上传大小（默认 32MiB）
	// router.MaxMultipartMemory = 8 << 20 //8MiB
	r.POST("/testUplad", func(c *gin.Context) { //路由接受POST请求 func接收gin的上下文、以c来表示
		// 单文件
		file, _ := c.FormFile("file")
		//c.SaveUploadedFile(file,"./"+file.Filename)

		//name := c.PostForm("name")
		// 文件流
		in,_ :=file.Open() // 打开文件 open
		defer in.Close() // 默认函数执行完关闭 close
		out, _ := os.Create("./"+file.Filename)
		io.Copy(out,in) // 本地写文件
	   //	defer out.Close()
		//c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s","文件"))
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s",file.Filename))
		c.File("./"+file.Filename) // 当前的path
		//c.JSON(200, gin.H{
		//	"msg":file,
		//	"name":name,
		//})
		log.Println(file.Filename)
	})
	r.Run(":8888")
}
