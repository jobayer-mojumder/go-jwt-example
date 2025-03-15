package routes

import (
	"go-jwt-project/internal/handlers"
	"go-jwt-project/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	//api version 1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "success",
				"message": "Rest API with JWT version 1",
			})
		})

		v1.POST("/login", handlers.Login)
		v1.POST("/register", handlers.Register)

		postGroupV1 := v1.Group("/posts")
		postGroupV1.Use(middlewares.JWTAuth())
		{
			postGroupV1.GET("/", handlers.GetPosts)
			postGroupV1.POST("/", handlers.CreatePost)
		}
	}

	//api version 2

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Route not found",
		})
	})

}
