package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mselser95/microservice-template/api/handlers"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.GET("/", handlers.GetUser)
}

func RegisterRoutes(r *gin.Engine) {
	RegisterUserRoutes(r)
	// Add more routes for other resources here
}
