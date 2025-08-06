package config

import (
	"log"

	"github.com/joho/godotenv"
)

/*
初始化加载环境变量功能(默认为 .env 文件)
*/
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
