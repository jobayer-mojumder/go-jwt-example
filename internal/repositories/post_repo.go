package repositories

import (
	"go-jwt-project/internal/models"

	"go-jwt-project/internal/database"
)

func CreatePost(post *models.Post) error {
	return database.DB.Create(post).Error
}

func GetPosts() ([]models.Post, error) {
	var posts []models.Post
	err := database.DB.Find(&posts).Error
	return posts, err
}
