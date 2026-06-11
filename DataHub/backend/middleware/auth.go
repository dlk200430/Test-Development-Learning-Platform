package middleware

import (
	"github.com/testdev-learning/DataHub/backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.ID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}