package main

import (
	"buble-master/database"
	"buble-master/routers"
	"fmt"
	"log"
)

func main() {
	//创建数据库 create database if not exists bubble default charset=utf8mb4;

	//初始化数据库连接
	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Databases connection failed :%v", err)
		return
	}
	// 设置路由
	r := routers.SetupRouter()

	//启动服务器
	port := ":10090"
	fmt.Printf("Server starting on port %s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Server failed to start:%v", err)
	}
}
