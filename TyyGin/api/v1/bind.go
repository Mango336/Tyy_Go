package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age"`
}

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

// querystring指的是浏览器中URL中 ? 后面携带的参数 【form数据】
func GetQueryString(c *gin.Context) {
	name := c.DefaultQuery("name", "defaultValue")
	age := c.Query("age")
	// age, ok := c.GetQuery("age") // ok返回是否获取到age字段
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

// func Get

// ShouldBind() 参数绑定 能基于请求自动提取JSON、form表单、QueryString的数据
// 并把值绑定到指定的结构体对象中
// ShouldBind()绑定顺序：1. 如果是GET请求=》 只使用Form绑定引擎（query）
// 2.如果是POST请求，首先检查content-type是否是json或xml 再用Form（form-data）

// GetShouleBindQuery ShouldBind()获取query string(path url)中的数据
// 例：43.142.18.112:9090/show/bind/query?name=Mango&age=25
func BindQueryString(c *gin.Context) {
	data := Person{}
	if err := c.ShouldBind(&data); err == nil {
		log.Println(data)
		c.JSON(http.StatusOK, gin.H{
			"name": data.Name,
			"age":  data.Age,
		})
	} else {
		log.Fatalln(err.Error())
	}
}

// BindJson 绑定JSON数据
func BindJson(c *gin.Context) {
	data := Person{}
	// 两个方法 ShouldBind ShouldBindJSON
	// if err := c.ShouldBind(&data); err == nil {
	if err := c.ShouldBindJSON(&data); err == nil {
		log.Println(data)
		c.JSON(http.StatusOK, gin.H{
			"name": data.Name,
			"age":  data.Age,
		})
	} else {
		log.Fatalln(err.Error())
	}
}

// BindForm 绑定JSON数据
func BindForm(c *gin.Context) {
	data := Person{}
	if err := c.ShouldBind(&data); err == nil {
		log.Println(data)
		c.JSON(http.StatusOK, gin.H{
			"name": data.Name,
			"age":  data.Age,
		})
	} else {
		log.Fatalln(err.Error())
	}
}
