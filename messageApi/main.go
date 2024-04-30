// Package main is used to start running the service
package main

import (
	"log"

	"messageApi/internal/database"
	"messageApi/internal/server"
	"messageApi/internal/service"
	"messageApi/internal/types"

	"github.com/caarlos0/env/v11"
)

// main creates the application modules and runs the server
func main() {
	cfg := types.Config{}

	// parse the environment variables
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	// initialize the data module
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// initialize the service module
	service, err := service.NewService(cfg, db)
	if err != nil {
		log.Fatal(err)
	}

	// initialize the server module
	srv, err := server.NewServer(cfg, service)
	if err != nil {
		log.Fatal(err)
	}

	// run the server
	srv.RunServer()
}
