package api

import (
	"employee-api/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheckAPI is a method to perform healthcheck of application
func HealthCheckAPI(c *gin.Context) {
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Employee API is not running. Check application logs")
	}
	defer scyllaClient.Close()
	c.JSON(http.StatusOK, gin.H{"message": "Employee API is up and running"})
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
