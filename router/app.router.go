package router

import (
	cityService "empire_origins_golang/city"
	"empire_origins_golang/middleware"
	userService "empire_origins_golang/user"

	"github.com/gin-gonic/gin"
)

/*
设置路由
*/
func SetupRouter() *gin.Engine {
	// 1. 创建一个gin实例, 用于处理请求
	router := gin.New()

	// 2. 设置路由, 当用户访问/user时, 调用service.GetUser函数
	{
		userRouter := router.Group("/user")
		{
			userRouter.GET("/info", userService.GetUser)
		}
		router.POST("/login", userService.Login)
		router.POST("/register", userService.Register)
		userRouter.Use(middleware.AuthMiddleware())
	}

	{
		cityRouter := router.Group("/city")
		{
			cityRouter.GET("/info", cityService.CityInfo)
		}
		cityRouter.Use(middleware.AuthMiddleware())
	}

	// 3. 返回路由实例
	return router
}
