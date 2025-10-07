package DAO

import (
	"Golang_task4/CONF"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB
var err error

// InitDB 初始化数据库连接
func InitDB() {
	// 加载配置
	CONF.LoadConfig()
	//gorm 链接数据库
	Conn, err = gorm.Open(mysql.Open(CONF.AppConfig.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	log.Println("数据库连接成功")
}

// 保持init函数用于向后兼容
func init() {
	InitDB()
}

// 执行表迁移
func TableAutoMigrate() {
	Conn.AutoMigrate(&User{})
	Conn.AutoMigrate(&Post{})
	Conn.AutoMigrate(&Comment{})
}
