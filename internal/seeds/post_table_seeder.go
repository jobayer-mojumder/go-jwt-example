package seeds

import (
	"go-jwt-project/internal/models"

	"gorm.io/gorm"
)

func PostsSeeder(db *gorm.DB) {
	posts := []models.Post{
		{
			Title:   "First Post",
			Content: "This is the first post",
			UserID:  1,
		},
		{
			Title:   "Second Post",
			Content: "This is the second post",
			UserID:  2,
		},
	}

	for _, post := range posts {
		db.Create(&post)
	}
}
