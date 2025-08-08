package router

import (
	"empire_origins_golang/middleware"
	"empire_origins_golang/service"

	"github.com/gin-gonic/gin"
)

/*
设置路由
*/
func SetupRouter() *gin.Engine {
	// 1. 创建一个gin实例, 用于处理请求
	router := gin.New()

	// 2. 设置路由, 当用户访问/user时, 调用service.GetUser函数
	userRouter := router.Group("/user")
	{
		userRouter.GET("/info", service.GetUser)
	}
	userRouter.Use(middleware.AuthMiddleware())

	router.POST("/login", service.Login)

	// 3. 返回路由实例
	return router
}
