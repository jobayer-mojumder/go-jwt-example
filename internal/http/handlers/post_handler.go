package handlers

import (
	"go-jwt-project/internal/database"
	"go-jwt-project/internal/http/requests"
	"go-jwt-project/internal/http/responses"
	"go-jwt-project/internal/logger"
	"go-jwt-project/internal/models"
	"go-jwt-project/internal/pkg/auth"
	"go-jwt-project/internal/repositories"
	"go-jwt-project/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func CreatePost(c *gin.Context) {
	var createPostRequest requests.CreatePostRequest

	if err := c.ShouldBindJSON(&createPostRequest); err != nil {

		logger.LogError(err, "Failed to bind JSON request")

		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post := models.Post{
		Title:   createPostRequest.Title,
		Content: createPostRequest.Content,
	}

	userID, err := auth.GetUserIDFromJWTGinContext(c)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Could not get user ID from JWT")
		return
	}

	post.UserID = userID

	if err := repositories.CreatePost(&post); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Could not create post")
		return
	}

	// Preload the user data associated with the post
	var createdPost models.Post
	if err := database.DB.Preload("User").First(&createdPost, post.ID).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Could not retrieve post data")
		return
	}

	var response responses.PostResponse
	if err := copier.Copy(&response, &createdPost); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Error copying post data")
		return
	}

	utils.SendSuccessResponse(c, http.StatusCreated, response)
}

func GetPosts(c *gin.Context) {
	posts, err := repositories.GetPosts(database.DB)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Could not retrieve posts")
		return
	}

	var response []responses.PostResponse
	for _, post := range posts {
		var postResponse responses.PostResponse
		if err := copier.Copy(&postResponse, &post); err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, "Error copying post data")
			return
		}

		response = append(response, postResponse)
	}

	utils.SendSuccessResponse(c, http.StatusOK, response)
}
