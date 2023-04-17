package model

import (
	"TyyGin/utils"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB // model包中可视
var errDB error

func InitDb() {
	// 数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName,
	)
	db, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Printf("Connecting mysql is error, error is %v...\n", errDB.Error())
		panic("Connecting mysql error...")
	}

	// 数据迁移
	db.AutoMigrate(&User{}) // User{} struct 相当于一个数据库表table
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Connecting mysql is error, error is %v...\n", err.Error())
		panic("Connecting mysql error...")
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
