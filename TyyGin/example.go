package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func example() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK...",
		})
	})
	LoadTemplate(r)
	r.Run(":9090") // listen :9090
}

// 加载Template html
func LoadTemplate(r *gin.Engine) {
	// 使用LoadHTMLGlob()或LoadHTMLFiles()方法进行HTML模板渲染
	r.LoadHTMLGlob("./templates/**/*") // 括号中是模板位置
	r.GET("/post/index", func(c *gin.Context) {
		// 第二个参数是对应html的name => 这里是posts/index.html是 posts中index.tmpl里面define的名称
		c.HTML(http.StatusOK, "post/index.html", gin.H{"title": "This is post/index"})
	})
	r.GET("/user/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{ // 第二个参数默认就是文件名
			"title":    "This is user/index",
			"username": "Mango",
			"age":      18,
		})
	})
	r.GET("/post/test", func(c *gin.Context) {
		// 第二个参数也可以直接匹配tmpl
		c.HTML(http.StatusOK, "index2.tmpl", gin.H{"title": "This is post/index2 for testing..."})
	})
}
