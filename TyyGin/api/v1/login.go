package v1

import (
	"TyyGin/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Load Login Web
func LoadLoginWeb(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}
func LoginPost(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	role := 0
	var data model.User = model.User{
		Username: username,
		Password: password,
		Role:     role,
	}
	fmt.Println(data)
	c.HTML(http.StatusOK, "login_post.tmpl", gin.H{
		"data": data,
	})
}
