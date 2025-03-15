package migrations

import (
	"go-jwt-project/internal/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// CreateUserTableMigration defines the migration for creating the 'users' table
func CreateUserTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "2025_03_15_001", // Migration ID (must be unique)
		Migrate: func(tx *gorm.DB) error {
			// Apply the migration (create users table)
			return tx.AutoMigrate(&models.User{})
		},
		Rollback: func(tx *gorm.DB) error {
			// Rollback the migration (drop users table)
			return tx.Migrator().DropTable(&models.User{})
		},
	}
}
