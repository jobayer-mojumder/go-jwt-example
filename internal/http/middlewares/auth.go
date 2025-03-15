package middlewares

import (
	"go-jwt-project/internal/pkg/auth"
	"go-jwt-project/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		_, claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		c.Set("id", claims["id"])
		c.Set("email", claims["email"])
		c.Next()
	}
}
