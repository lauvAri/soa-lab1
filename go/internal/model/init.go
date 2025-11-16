package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"materials-service/internal/config"
)

var DB *gorm.DB

func InitDB() error {
	dsn := config.GetDSN() // 从配置获取DSN
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 配置连接池
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	
	// 自动迁移（可选）
	// DB.AutoMigrate(&Material{}, &MaterialType{})
	return nil
}