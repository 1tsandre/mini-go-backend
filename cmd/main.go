package main

import (
	"fmt"

	"github.com/1tsandre/mini-go-backend/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)
}