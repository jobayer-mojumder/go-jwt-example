package repositories

import (
	"go-jwt-project/internal/models"

	"go-jwt-project/internal/database"

	"gorm.io/gorm"
)

func CreatePost(post *models.Post) error {

	if err := database.DB.Create(post).Error; err != nil {
		return err
	}

	return nil
}

func GetPosts(db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	if err := db.Preload("User").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}
