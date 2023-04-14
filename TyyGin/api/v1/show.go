package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowName(c *gin.Context) {
	c.HTML(http.StatusOK, "showname.tmpl", gin.H{
		"title": "子路由1--第一个网页",
		"name":  "Mango",
	})
}
func ShowInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "YLM",
		"sex":  "male",
		"age":  25,
	})
}
