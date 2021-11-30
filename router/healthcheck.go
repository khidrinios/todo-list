package router

import "github.com/gin-gonic/gin"

func HealthCheckInitRoute(router *gin.Engine) {
	router.GET("/", healthCheck)
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"application": "todo",
		"success":     true,
		"response":    "You are now under Todo's shade. https://www.youtube.com/watch?v=xos2MnVxe-c",
	})
}
