// model/user.go
package model

// import "gorm.io/gorm"

type User struct {
	// gorm.Model
	// 注意tag中 key冒号后不能空格
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Role     int    `form:"role" json:"role"`
}
