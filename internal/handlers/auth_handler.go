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
	// get the email and password from the request body
	email := c.PostForm("email")
	password := c.PostForm("password")

	// check if the email and password are valid and not empty
	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	// get the user by email
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// check if the password is correct
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

	// return a success message if the user is logged in and create a JWT token
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User logged in", "token": token})
}

func Register(c *gin.Context) {
	// get all the data from the request body
	email := c.PostForm("email")
	password := c.PostForm("password")
	name := c.PostForm("name")

	// check if the email and password are valid and not empty
	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	// check if the user already exists
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

	// return a success message if the user is created and create a JWT token
	token, err := auth.GenerateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "token": token})

}
