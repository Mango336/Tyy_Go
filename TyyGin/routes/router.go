package routes

import (
	v1 "TyyGin/api/v1"
	"TyyGin/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.LoadHTMLGlob("./web/*") // 括号中是模板位置
	// 使用中间件
	// Recovery 中间件会recover任何panic 如果有panic的话，会写入500
	r.Use(utils.HelloWeb(), gin.Recovery())

	// RESTful
	// GET 获取资源；POST 新建资源；PUT 更新资源；DELETE 删除资源；
	// 展示测试模块
	rp1 := r.Group("/show")
	{
		rp1.GET("/showName", v1.ShowName)
		rp1.GET("/showInfo", v1.ShowInfo)
	}
	// 登录模块路由
	rp2 := r.Group("/login")
	{
		rp2.GET("/", v1.LoadLoginWeb)
		rp2.POST("/", v1.LoginPost)
	}
	r.Run(":9090")
}
