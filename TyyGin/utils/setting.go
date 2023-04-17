// 加载配置文件
package utils

import (
	"log"

	ini "gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Printf("Error is: %v\n", err.Error())
		panic("Config path is wrong, can not read this path...")
	}
	LoadServer(file)
	LoadDataBase(file)
}

// 加载Server配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("Server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("Server").Key("HttpPort").MustString(":9090")
}

// 加载DataBase配置
func LoadDataBase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("dase618")
	DbName = file.Section("database").Key("DbName").MustString("root")
}
