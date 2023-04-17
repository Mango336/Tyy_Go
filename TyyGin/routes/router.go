package routes

import (
	v1 "TyyGin/api/v1"
	"TyyGin/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()             // 声明一个默认路由
	r.MaxMultipartMemory = 8 << 20 // 甚至内存限制为8MB

	r.LoadHTMLGlob("./web/*") // 括号中是模板位置
	// 使用中间件
	// Recovery 中间件会recover任何panic 如果有panic的话，会写入500
	r.Use(utils.HelloWeb(), gin.Recovery()) // 全局中间件

	// RESTful
	// GET 获取资源；POST 新建资源；PUT 更新资源；DELETE 删除资源；
	// 展示测试模块路由
	rpShow := r.Group("/show")
	{
		rpShow.GET("/showName", v1.ShowName)
		rpShow.GET("/showInfo", v1.ShowInfo)
		subRpShow := rpShow.Group("/bind")
		{
			subRpShow.GET("/:name/:age", v1.GetPathUrl)
			subRpShow.GET("/querystring", v1.GetQueryString)
			subRpShow.GET("/query", v1.BindQueryString)
			subRpShow.POST("/json", v1.BindJson)
			subRpShow.POST("/form", v1.BindForm)

		}
	}
	// 登录模块路由
	rpLogin := r.Group("/login")
	{
		rpLogin.GET("/", v1.LoadLoginWeb)
		rpLogin.POST("/", v1.Login)
		rpLogin.POST("/signup", v1.SignUp)
		rpLogin.GET("/select", v1.SelectUser)
	}

	// 上传文件路由
	rpUpload := r.Group("/upload")

	{
		rpUpload.POST("/", v1.UpLoadFile)
		rpUpload.POST("/files", v1.UpLoadFiles)
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "Welcome to server01...",
		})
	})
	return e
}

func Router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "Welcome to server02...",
		})
	})
	return e
}
