package api

import (
	"github.com/gin-gonic/gin"
	public "github.com/mselser95/microservice-template/api/routes"
	v1 "github.com/mselser95/microservice-template/api/routes/v1"
	"github.com/mselser95/microservice-template/internal/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"time"
)

// Custom middleware to log requests using Zap
func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request details after response is sent
		logger.Info("Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		)
	}
}

// Recovery middleware with Zap logging
func ZapRecovery(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log the panic with Zap
				logger.Error("Panic recovered",
					zap.Any("error", r),
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
				)
				c.AbortWithStatus(500) // Optional: Return 500 Internal Server Error
			}
		}()
		// Continue with the request
		c.Next()
	}
}

func NewApiServer(
	cfg config.Config,
	logger *zap.Logger,
) {

	// Create a new Gin router without the default logger and recovery middleware
	r := gin.New()

	// Use the custom Zap logger middleware
	r.Use(ZapLogger(logger))

	// Optionally, use a custom recovery middleware (with Zap logging)
	r.Use(ZapRecovery(logger))

	// Use Zap as the logger middleware
	r.Use(ZapLogger(logger))

	// Register Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Register all v1 routes (which are authenticated)
	v1.RegisterV1Routes(r)
	public.RegisterRoutes(r)

	// Get the port from the environment variable or default to 8080
	port := "8080"

	logger.Info("starting-server", zap.String("port", port))

	// Start the server on the specified port
	err := r.Run(":" + port)
	if err != nil {
		logger.Error("error-starting-server", zap.Error(err))
	}
}
