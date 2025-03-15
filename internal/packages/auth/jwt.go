package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    userID,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return token, *claims, nil
	} else {
		return nil, nil, err
	}
}

func GetUserIDFromJWTGinContext(c *gin.Context) (uint, error) {
	id, exists := c.Get("id")
	if !exists {
		return 0, fmt.Errorf("user ID not found in context")
	}

	idValue, ok := id.(float64)
	if !ok {
		return 0, fmt.Errorf("invalid user ID type")
	}

	return uint(idValue), nil
}

func GetEmailFromJWTGinContext(c *gin.Context) (string, error) {
	email, exists := c.Get("email")
	if !exists {
		return "", fmt.Errorf("email not found in context")
	}

	emailValue, ok := email.(string)
	if !ok {
		return "", fmt.Errorf("invalid email type")
	}

	return emailValue, nil
}
