package middlewares

import (
	"fmt"
	"go-jwt-project/internal/pkg/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		_, claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("id", claims["id"])
		c.Set("email", claims["email"])

		fmt.Printf("User ID: %v\n", claims["id"].(float64)) // Prints the actual user ID
		fmt.Printf("Email: %v\n", claims["email"].(string)) // Prints the actual email
		c.Next()
	}
}
