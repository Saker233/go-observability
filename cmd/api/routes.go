package main

import (
	"github.com/gin-gonic/gin"

	appmetrics "Saker233/go-observability/prometheus"
)

func routes() *gin.Engine {
	router := gin.Default()

	router.Use(PrometheusMiddleware())
	router.GET("/health", getHealth)
	router.GET("/work", doWork)
	router.GET("/metrics", gin.WrapH(appmetrics.Handler()))

	return router
}
