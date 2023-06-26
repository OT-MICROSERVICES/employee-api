package api

import (
	"employee-api/client"
	"employee-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary HealthCheckAPI is a method to perform healthcheck of application
// @Schemes http
// @Description Do healthcheck
// @Tags healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} model.HealthCheck
// @Router /health [get]
// HealthCheckAPI is a method to perform healthcheck of application
func HealthCheckAPI(c *gin.Context) {
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		data := model.HealthCheck{
			Message: "Employee API is not running. Check application logs",
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	defer scyllaClient.Close()
	data := model.HealthCheck{
		Message: "Employee API is up and running",
	}
	c.JSON(http.StatusOK, data)
}

// @Summary DetailedHealthCheckAPI is a method to perform detailed healthcheck of application
// @Schemes http
// @Description Do detailed healthcheck
// @Tags healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} model.DetailedHealthCheck
// @Router /health/detail [get]
// DetailedHealthCheckAPI is a method to perform detailed healthcheck of application
func DetailedHealthCheckAPI(c *gin.Context) {
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		data := model.DetailedHealthCheck{
			Message:          "Employee API is not running. Check application logs",
			ScyllaDB:    "down",
			EmployeeAPI: "down",
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	defer scyllaClient.Close()
	data := model.DetailedHealthCheck{
		Message:          "Employee API is up and running",
		ScyllaDB:    "up",
		EmployeeAPI: "up",
	}
	c.JSON(http.StatusOK, data)
}

