package main

import (
	"log"

	"github.com/1tsandre/mini-go-backend/internal/app"
	"github.com/1tsandre/mini-go-backend/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	application.Start()
}
