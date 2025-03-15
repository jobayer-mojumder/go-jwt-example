package routes

import (
	"go-jwt-project/internal/handlers"
	"go-jwt-project/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)

	postGroup := router.Group("/posts")
	postGroup.Use(middlewares.JWTAuth())
	{
		postGroup.GET("/", handlers.GetPosts)
		postGroup.POST("/", handlers.CreatePost)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Route not found",
		})
	})

}
