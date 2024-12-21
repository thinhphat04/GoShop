package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// LoggerMiddleware logs the details of each request
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Tiếp tục xử lý request
		c.Next()

		// Ghi log sau khi xử lý
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()
		log.Printf("[REQUEST] %s %s | Status: %d | Duration: %v",
			c.Request.Method, c.Request.URL.Path, statusCode, duration)
	}
}
