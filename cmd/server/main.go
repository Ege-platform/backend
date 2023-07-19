package main

import (
	"ege_platform/internal/api/router"
	"ege_platform/internal/config"
	"ege_platform/internal/logging"
	"ege_platform/internal/pb"
	"fmt"
	"os"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Error while creating config: %v", err)
		os.Exit(-1)
	}

	err = logging.Init(cfg)
	if err != nil {
		fmt.Printf("Error while initializing logging: %v", err)
		os.Exit(-1)
	}

	p := pb.NewPB(cfg)

	router := router.NewRouter(p, cfg)
	router.SetupRoutes()

	if err := p.Run(); err != nil {
		logging.Log.Fatalf("Error starting server: %v", err)
	}
}
