package middleware

import (
	"empire_origins_golang/model"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Authorization 头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		// 2. 切分 authHeader 为 "Bearer " 和 token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		// 3. 验证 token
		token_string := parts[1]
		token, err := jwt.ParseWithClaims(token_string, model.UserInfo{}, func(token *jwt.Token) (any, error) {
			// 验证 token 的签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			// 返回 token 的密钥
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.JSON(403, gin.H{"message": "Forbidden"})
			c.Abort()
			return
		}
		// 4. 获取用户信息
		userInfo := token.Claims.(model.UserInfo)
		// 5. 将用户信息添加到上下文
		c.Set("userInfo", userInfo)
		// 6. 继续处理请求
		c.Next()
	}
}
