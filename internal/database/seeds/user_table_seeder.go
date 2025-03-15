package seeds

import (
	"go-jwt-project/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UsersSeeder(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)

	if count > 0 {
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	users := []models.User{
		{
			Name:     "Sample User",
			Email:    "test@gmail.com",
			Password: hashedPassword,
		},
		{
			Name:     "John Doe",
			Email:    "doe@gmail.com",
			Password: hashedPassword,
		},
	}

	for _, user := range users {
		db.Create(&user)
	}
}
