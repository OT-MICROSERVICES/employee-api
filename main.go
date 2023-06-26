package main

import (
	"employee-api/routes"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	v1 := router.Group("/v1")
	routes.CreateRouterForEmployee(v1)
	router.Run(":8080")
}
