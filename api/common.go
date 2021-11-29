package api

import "github.com/gin-gonic/gin"

func HandleResponse(c *gin.Context, statusCode int, response interface{}) {
	c.JSON(statusCode, response)
}

func HandleResponseError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}
