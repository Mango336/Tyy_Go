// model/user.go
package model

import (
	"TyyGin/utils/errmsg"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// 注意tag中 key冒号后不能空格
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Role     int    `form:"role" json:"role"`
}

// 创建用户
func CreateUser(user *User) int {
	result := db.Create(user) // 使用数据的指针来创建
	if result.Error != nil {  // 返回的Error
		log.Println("Create user fail...")
		return errmsg.ERROR
	}
	log.Println("插入的用户ID为：", user.ID)
	log.Println("插入记录的条数为：", result.RowsAffected)
	return errmsg.SUCCESS
}

// 查询用户
func RetrieveUser(name, pwd string) int {
	user := User{}
	db.Where("username = ? AND password = ?", name, pwd).Find(&user)
	log.Println(user)

	if user.ID < 1 { // ID从1开始 小于1的话说明没有找到该记录
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}
