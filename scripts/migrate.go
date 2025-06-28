package main

import (
	"log"

	"golang/internal/config"
	"golang/internal/database"
	"golang/internal/models"
)

func main() {
	// Carregar configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Conectar ao banco de dados
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close(db)

	// Executar migrações
	log.Println("Running database migrations...")

	// Adicione seus modelos aqui para migração
	err = db.AutoMigrate(
		&models.User{},
		// Adicione mais modelos conforme necessário
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed successfully!")
}
