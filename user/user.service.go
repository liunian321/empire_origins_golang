package service

import (
	"empire_origins_golang/config"
	"empire_origins_golang/db"
	"empire_origins_golang/dto"
	"empire_origins_golang/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

// @Summary 登录
// @Description 登录
// @Tags users
// @Accept  json
// @Produce  json
// @Param   loginUser     body    model.LoginUser     true  "登录用户"
func Login(c *gin.Context) {
	// 1. 校验请求参数
	loginUser := dto.LoginUser{}
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 2. 查询用户,检验是否存在
	user, err := gorm.G[model.User](db.DB).Select("id, password, \"order\", username").Where("username = ?", loginUser.Username).First(c)
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
		mapElementIds[i] = city.MapElementID
	}

	if len(mapElementIds) == 0 {
		c.JSON(500, gin.H{"error": "User has no cities"})
		return
	}

	// 6. 生成 token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"order":    user.Order,
		"cities":   mapElementIds,
	}).SignedString(config.JwtSecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": dto.AddUserResponse{
		UserId:       user.ID,
		AccessToken:  token,
		MapElementId: mapElementIds[0],
	}})
}

func Register(c *gin.Context) {
	registerUser := dto.LoginUser{}
	if err := c.ShouldBindJSON(&registerUser); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate password hash"})
		return
	}
	user := model.User{}

	user.Password = string(password)
	user.Username = registerUser.Username
	user.IsEnabled = true
	user.Order = 0
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = gorm.G[model.User](db.DB).Create(c, &user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"order":    user.Order,
	}).SignedString(config.JwtSecret)
	c.JSON(200, gin.H{"token": token})
}
