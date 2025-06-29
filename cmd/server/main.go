package main

import (
	"log"
	"os"

	"golang/internal/api"
	"golang/internal/config"
	"golang/internal/database"
	"golang/internal/middleware"
)

func main() {
	// Carregar configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Inicializar logger
	logger := middleware.NewLogger()

	// Conectar ao banco de dados
	db, err := database.Connect(cfg.Database)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	// Criar servidor HTTP
	server := api.NewServer(cfg, db, logger)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Infof("Starting server on port %s", port)
	if err := server.Start(":" + port); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
