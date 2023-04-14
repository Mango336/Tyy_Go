package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPathUrl 获取path（url）中携带的参数: 返回的都是string类型
func GetPathUrl(c *gin.Context) {
	// url为43.142.18.112:9090/show/get/Mango/25
	// 获取到username= Mango age=25
	name := c.Param("name") // Param(key string) key为接口中定义的名
	age, _ := strconv.Atoi(c.Param("age"))
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

// GetInfoBind 参数绑定 ShouldBind()能基于请求自动提取JSON、form表单、QueryString的数据 并把值绑定到指定的结构体对象中
func GetInfoBind() {

}
