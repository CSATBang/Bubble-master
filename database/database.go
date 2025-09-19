package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局数据库连接变量
var DB *gorm.DB

// ConnectDB 初始化数据库连接
func ConnectDB() error {
	var err error
	// DSN (Data Source Name) 格式: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := "root:xxxxxx@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect databases:%w", err)
	}
	fmt.Println("DB链接成功!")
	return nil
}

// GetDB 获取数据库实例（可选）
func GetDB() *gorm.DB {
	return DB
}
