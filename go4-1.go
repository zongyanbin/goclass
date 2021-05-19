/**
多文件上传
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
//c.SaveUploadedFile(file,"./"+file.Filename)
func main(){
	 r := gin.Default() // 路由启动gin服务 携带基数中间件启动
	 r.POST("/testUpload",postFile)
	 r.Run(":8888")
}

// 路由接收POST请求 func里接收gin上下文 以c来表示
func postFile(c *gin.Context) {
	form,err :=c.MultipartForm()
	files  := form.File["file[]"]
	fmt.Println(files)
	//错误处理
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	return
	}
	for _, files := range files{
		log.Println(files.Filename)
		// 上传文件到指定的路径
		c.SaveUploadedFile(files,"./upload/"+files.Filename)
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s",files.Filename))
		c.File("./"+files.Filename) // 当前的path
		//c.JSON(200, gin.H{
		//	"msg":files,
		//	"name":123,
		//})

	}
	c.String(http.StatusOK,fmt.Sprintf("%d files uploaded!",files))
}

// 检测
//func check(err){
//	if err != nil{
//		panic(err)
//	}
//}