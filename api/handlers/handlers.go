package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser handles GET requests for retrieving a user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	// Fetch the user from the database or service
	user := map[string]interface{}{
		"id":    id,
		"name":  "John Doe",
		"email": "johndoe@example.com",
	}
	c.JSON(http.StatusOK, user)
}
