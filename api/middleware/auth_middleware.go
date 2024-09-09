package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware checks for a valid Authorization token in the header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// Example: Validate the token (this is a simple mock check)
		if token != "valid_token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// If the token is valid, proceed to the next handler
		c.Next()
	}
}
