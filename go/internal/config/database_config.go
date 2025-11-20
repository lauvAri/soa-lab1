package config

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("无法加载 .env 文件")
	}
}

func GetDSN() string {
	raw := os.Getenv("DB_DSN")
	if raw == "" {
		return raw
	}

	cfg, err := mysql.ParseDSN(raw)
	if err != nil {
		fmt.Println("解析 DSN 失败，使用原始字符串:", err)
		return raw
	}
	cfg.ParseTime = true
	if cfg.Params == nil {
		cfg.Params = map[string]string{}
	}
	if _, ok := cfg.Params["charset"]; !ok {
		cfg.Params["charset"] = "utf8mb4"
	}
	return cfg.FormatDSN()
}
