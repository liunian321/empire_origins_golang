package service

import (
	"empire_origins_golang/db"
	"empire_origins_golang/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	// 1. 校验请求参数
	loginUser := model.LoginUser{}
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 2. 查询用户,检验是否存在
	user, err := gorm.G[model.User](db.DB).Select("id, password, order, username").Where("username = ?", loginUser.Username).First(c)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// 3. 检验密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	// 4. 查询用户的城市信息
	var cities []model.City
	cities, err = gorm.G[model.City](db.DB).Select("map_element_id").Where("owner_id = ?", user.ID).Find(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to query user cities"})
		return
	}

	// 5. 生成 token
	mapElementIds := make([]string, len(cities))
	for i, city := range cities {
		mapElementIds[i] = city.MapElementId
	}

	// 检查用户是否有城市
	if len(mapElementIds) == 0 {
		c.JSON(500, gin.H{"error": "User has no cities"})
		return
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"order":    user.Order,
		"cities":   mapElementIds,
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
