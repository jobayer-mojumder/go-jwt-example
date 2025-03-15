package seeds

import (
	"gorm.io/gorm"
)

// run all seeds
func Run(db *gorm.DB) {

	// list of all seeders in array and then loop through them
	var seeders = []func(*gorm.DB){
		UsersSeeder,
		PostsSeeder,
	}

	for _, seeder := range seeders {
		seeder(db)
	}
}
