package handlers

import (
	"go-jwt-project/internal/models"
	"go-jwt-project/internal/pkg/auth"
	"go-jwt-project/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// create a JWT token
	token, err := auth.GenerateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User logged in", "token": token})
}

func Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	name := c.PostForm("name")

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	if _, err := repositories.GetUserByEmail(email); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// hash the password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// create the user with hashed password
	user := models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
	repositories.CreateUser(&user)

	token, err := auth.GenerateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "token": token})
}
