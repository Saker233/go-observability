package main

import (
	"github.com/gin-gonic/gin"
)

func routes() *gin.Engine {
	router := gin.Default()

	router.GET("/health", getHealth)
	router.GET("/work", doWork)

	return router
}
