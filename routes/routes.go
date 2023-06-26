package routes

import (
	"employee-api/api"
	"github.com/gin-gonic/gin"
)

// CreateRouterForEmployee is a method for generate routes
func CreateRouterForEmployee(routerGroup *gin.RouterGroup) {
	employee := routerGroup.Group("/employee")

	employee.GET("/health", api.HealthCheckAPI)
}
