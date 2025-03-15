package repositories

import (
	"go-jwt-project/internal/database"
	"go-jwt-project/internal/models"
)

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return user, err
}

func CreateUser(user *models.User) error {
	err := database.DB.Create(&user).Error
	return err
}

func UpdateUser(user *models.User) error {
	err := database.DB.Save(&user).Error
	return err
}

func DeleteUser(user *models.User) error {
	err := database.DB.Delete(&user).Error
	return err
}

func CheckUserPassword(password, hash string) bool {
	return password == hash
}
