package main

import (
	"TyyGo/utils"
	"fmt"
)

func init() {
	fmt.Println("Init in main package...")
}

func main() {
	fmt.Println("This is a main package...") // 主函数
	// utils.GolangBase()                       // golang基础
	utils.GolangSync() // golang并发编程
}
