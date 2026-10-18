package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Healthy",
	})
}

func doWork(c *gin.Context) {
	// simulate work
	time.Sleep(500 * time.Millisecond)

	c.JSON(200, gin.H{
		"status": "work done!",
	})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
