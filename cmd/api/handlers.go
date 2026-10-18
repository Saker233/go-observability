package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func getHealth(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message": "Healthy",
	})
}


func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}