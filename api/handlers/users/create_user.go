package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser handles POST requests for creating a new user
func CreateUser(c *gin.Context) {
	var newUser map[string]interface{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add logic to save the new user
	c.JSON(http.StatusCreated, gin.H{"status": "User created", "user": newUser})
}

func createUser() {
	// This function is not exported and cannot be accessed outside of this package
	// for testing purposes...
}
