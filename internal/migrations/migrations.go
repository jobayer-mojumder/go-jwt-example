package migrations

import (
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	// Define migrations
	migrationsList := []*gormigrate.Migration{
		CreateUserTableMigration(),
		CreatePostTableMigration(),
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, migrationsList)

	// Run migrations (up)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)

		// Rollback failed migrations
		if err := m.RollbackLast(); err != nil {
			log.Fatalf("Could not rollback last migration: %v", err)
		}
	}
	fmt.Println("Migrations applied successfully")

	return m.Migrate()
}
