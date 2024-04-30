package server

import (
	"messageApi/internal/service"

	"github.com/gin-gonic/gin"
)

// ServiceMiddleware adds the service to the context for use in request processing
func ServiceMiddleware(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("service", service)
		c.Next()
	}
}
