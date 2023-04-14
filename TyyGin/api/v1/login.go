package v1

import (
	"TyyGin/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoadLoginWeb 加载登录页面
func LoadLoginWeb(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

// LoginPost 获取登录表单的数据
func LoginPost(c *gin.Context) {
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
		log.Println(data, data.Username, data.Password, data.Role)
		c.HTML(http.StatusOK, "login_post.tmpl", gin.H{
			"data": data,
		})
	} else {
		log.Fatal("Could not bind form info...", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

}
