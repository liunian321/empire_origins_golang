package service

import (
	"empire_origins_golang/db"
	"empire_origins_golang/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary 获取用户信息
// @Description 根据 ID 返回用户详情
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id     path    int     true  "用户ID"
// @Success 200 {object} model.User
// @Failure 400 {string} string "Bad request"
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Query("id")
	// 1. 检查请求参数
	if id == "" {
		c.JSON(400, gin.H{"message": "id is required"})
		return
	}
	// 2. 查询用户信息 (新语法 Generics API, 增删改查前面基本都一样的写法,就是最后的方法名不一样.)
	user, err := gorm.G[model.User](db.DB).Where("id = ?", id).First(c)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}
	// 3. 返回用户信息(自动序列化成 json 格式)
	c.JSON(200, user)
}
