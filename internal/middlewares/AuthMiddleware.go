package middlewares

import (
	"net/http"
	"strings"

	"github.com/alerdn/rest-go/internal/helpers/auth"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if !strings.HasPrefix(authorization, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Acesso não autorizado"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authorization, "Bearer ")
		claims, err := auth.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Acesso não autorizado"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)
		c.Next()
	}
}
