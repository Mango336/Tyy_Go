package v1

import (
	"TyyGin/model"
	"TyyGin/utils/errmsg"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoadLoginWeb 加载登录页面
func LoadLoginWeb(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

// LoginPost 获取登录表单的数据
func Login(c *gin.Context) {
	// 方法一：
	// username, _ := c.GetPostForm("username")
	// password, _ := c.GetPostForm("password")
	// role := 0
	// var data model.User = model.User{
	// 	Username: username,
	// 	Password: password,
	// 	Role:     role,
	// }
	// log.Println(data)
	// c.HTML(http.StatusOK, "login_post.tmpl", gin.H{
	// 	"data": data,
	// })

	// 方法二：
	var data model.User
	if err := c.ShouldBind(&data); err == nil {
		// log.Println(data, data.Username, data.Password, data.Role)
		c.HTML(http.StatusOK, "login_post.tmpl", gin.H{
			"data": data,
		})
	} else {
		log.Println("Could not bind form info...", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

// 注册用户
func SignUp(c *gin.Context) {
	// 1. 绑定数据
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": errmsg.ERROR})
	}
	// 2. 数据库操作
	code := model.CreateUser(&user)
	// 3. 后续数据库操作
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"massage": errmsg.GetErrMsg(code),
	})
}

// 查询用户
func SelectUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil { // 绑定数据
		c.JSON(http.StatusOK, gin.H{"status": errmsg.ERROR})
	}
	log.Println(user)
	code := model.RetrieveUser(user.Username, user.Password) // 查询数据
	fmt.Println(code)
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": errmsg.GetErrMsg(code)})
}
