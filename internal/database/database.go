package database

import (
	"go-jwt-project/internal/helpers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := helpers.GetEnv("DB_USER", "root") + ":" + helpers.GetEnv("DB_PASSWORD", "root") + "@tcp(127.0.0.1:3306)/" + helpers.GetEnv("DB_NAME", "go_jwt_project") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db
}
