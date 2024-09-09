package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mselser95/microservice-template/api/handlers/users"
	"github.com/mselser95/microservice-template/api/middleware"
)

func RegisterV1Routes(r *gin.Engine) {
	// Define a new v1 group that requires authentication
	v1Group := r.Group("/v1", middleware.AuthMiddleware())
	{
		userRoutes := v1Group.Group("/users")
		{
			userRoutes.GET("/:id", users.GetUser)  // Requires authentication
			userRoutes.POST("/", users.CreateUser) // Requires authentication
		}
	}
}
