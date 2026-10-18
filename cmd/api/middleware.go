package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	appmetrics "Saker233/go-observability/prometheus"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		appmetrics.WorkInFlight.Inc()
		defer appmetrics.WorkInFlight.Dec()

		c.Next()

		status := strconv.Itoa(c.Writer.Status())

		appmetrics.WorkRequests.
			WithLabelValues(status).
			Inc()

		appmetrics.WorkDuration.
			Observe(time.Since(start).Seconds())
	}
}
