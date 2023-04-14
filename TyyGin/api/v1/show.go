package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowName 简单的展示 向前端响应一个html页面
func ShowName(c *gin.Context) {
	c.HTML(http.StatusOK, "showname.tmpl", gin.H{
		"title": "子路由1--第一个网页",
		"name":  "Mango",
	})
}

// ShowInfo 向前端发送简单的json数据
func ShowInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "YLM",
		"sex":  "male",
		"age":  25,
	})
}
