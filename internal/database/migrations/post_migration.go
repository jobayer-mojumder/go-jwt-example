package migrations

import (
	"go-jwt-project/internal/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreatePostTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "2025_03_15_002",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Post{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&models.Post{})
		},
	}
}
