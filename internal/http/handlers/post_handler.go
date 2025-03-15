package handlers

import (
	"go-jwt-project/internal/database"
	"go-jwt-project/internal/http/requests"
	"go-jwt-project/internal/http/responses"
	"go-jwt-project/internal/models"
	"go-jwt-project/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var createPostRequest requests.CreatePostRequest

	if err := c.ShouldBindJSON(&createPostRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	post.Title = createPostRequest.Title
	post.Content = createPostRequest.Content

	userID, _ := c.Get("userID")
	post.UserID = userID.(uint)

	if err := repositories.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func GetPosts(c *gin.Context) {
	posts, err := repositories.GetPosts(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve posts"})
		return
	}

	var response []responses.PostResponse
	for _, post := range posts {
		response = append(response, responses.PostResponse{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			User: responses.UserResponse{
				ID:    post.User.ID,
				Name:  post.User.Name,
				Email: post.User.Email,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}
