package utils

import "github.com/gin-gonic/gin"

func SendErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}

func SendSuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}
