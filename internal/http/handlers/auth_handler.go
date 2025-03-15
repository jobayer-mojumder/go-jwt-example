package handlers

import (
	"go-jwt-project/internal/http/requests"
	"go-jwt-project/internal/logger"
	"go-jwt-project/internal/models"
	"go-jwt-project/internal/pkg/auth"
	"go-jwt-project/internal/repositories"
	"go-jwt-project/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	var loginRequest requests.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		logger.LogError(err, "Failed to bind JSON request")
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	email := loginRequest.Email
	password := loginRequest.Password

	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "User not found")
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Invalid password")
		return
	}

	token, err := auth.GenerateJWT(user.ID, user.Email)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create JWT token")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, gin.H{"message": "User logged in", "token": token})
}

func Register(c *gin.Context) {

	var registerRequest requests.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		logger.LogError(err, "Failed to bind JSON request")
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	name := registerRequest.Name
	email := registerRequest.Email
	password := registerRequest.Password

	if _, err := repositories.GetUserByEmail(email); err == nil {
		utils.SendErrorResponse(c, http.StatusConflict, "User already exists")
		return
	}

	// hash the password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
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
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create JWT token")
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, gin.H{"message": "User created", "token": token})
}
