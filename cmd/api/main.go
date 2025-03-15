package main

import (
	"go-jwt-project/internal/config"
	"go-jwt-project/internal/database"
	"go-jwt-project/internal/database/migrations"
	"go-jwt-project/internal/database/seeds"
	"go-jwt-project/internal/routes"
	"go-jwt-project/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvVariables()

	database.ConnectDB()

	// run all migrations
	err := migrations.RunMigrations(database.DB)
	if err != nil {
		log.Fatal(err)
	}

	// run all seeds
	seeds.Run(database.DB)

	port := utils.GetEnv("PORT", "8080")

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":" + port)

	// print the port the server is running on to the console
	println("Server running on port: " + port)
}
