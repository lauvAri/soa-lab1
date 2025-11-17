package config

import (
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

func LoadEnv() {
	err := godotenv.Load(".env") // 从 .env 文件加载配置
	if err != nil {
		fmt.Println("无法加载 .env 文件")
	}
}

func GetDSN() string {
	msg := os.Getenv("DB_DSN") 

	fmt.Println("DB_DSN:",msg)

	return os.Getenv("DB_DSN") 
}