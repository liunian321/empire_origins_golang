package main

import (
	"os"

	"empire_origins_golang/config"
	"empire_origins_golang/db"
	"empire_origins_golang/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// 1. 初始化环境变量环境
	config.LoadEnv()

	// 2. 初始化数据库连接
	db.Init_Postgres()

	// 3. 初始化 redis 连接
	db.Init_Redis()

	// 4. 从环境中获取端口号
	port := os.Getenv("PORT")
	if port == "" {
		port = "6001"
	}

	// 5. 设置 swagger 文档地址
	url := ginSwagger.URL("http://localhost:" + port + "/swagger/doc.json")
	router.SetupRouter().GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 6. 启动服务器, 并且监听端口
	router.SetupRouter().Run(":" + port)
}
