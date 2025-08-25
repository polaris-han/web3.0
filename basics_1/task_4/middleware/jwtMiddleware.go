package middleware

import (
	"fmt"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Authorization头获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 检查token格式（Bearer <token>）
		const bearerPrefix = "Bearer "
		if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		// 提取token字符串
		tokenString := authHeader[len(bearerPrefix):]

		// 解析token
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return service.JwtSecret, nil
		})

		// 验证token
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		fmt.Println("User ID:", claims["id"])
		fmt.Println("Username:", claims["username"])

		// 将解析后的用户信息存储到上下文
		c.Set("userID", claims["id"])
		c.Set("username", claims["username"])
		c.Set("claims", claims) // 存储完整的声明

		// 继续处理请求
		c.Next()
	}
}
