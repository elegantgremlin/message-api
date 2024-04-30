package server

import (
	"github.com/gin-gonic/gin"

	"messageApi/internal/service"
	"messageApi/internal/types"
)

// Server is the interface for the server module of the application
type Server interface {
	RunServer()
}

// server is the implementation of the server module
type server struct {
	service service.Service
	api     *gin.Engine
}

// NewServer creates an instance of the server module
func NewServer(cfg types.Config, service service.Service) (Server, error) {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	v1Group := r.Group("/v1")
	addV1Routes(v1Group, service)

	return &server{service, r}, nil
}

// RunServer starts running the server
func (s *server) RunServer() {
	s.api.Run()
}

// errorAsJSON converts an error message to a JSON payload for return in the response body
func errorAsJSON(msg string) map[string]string {
	return map[string]string{"error": msg}
}
