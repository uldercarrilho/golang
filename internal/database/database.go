package database

import (
	"fmt"

	"golang/internal/config"
	"golang/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect estabelece conexão com o banco de dados.
func Connect(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err) //nolint:wrapcheck
	}

	// Auto-migrate models
	if err := autoMigrate(db); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate: %w", err) //nolint:wrapcheck
	}

	return db, nil
}

// autoMigrate executa as migrações automáticas dos modelos.
func autoMigrate(db *gorm.DB) error {
	// Adicione seus modelos aqui para auto-migração
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("failed to auto-migrate user model: %w", err) //nolint:wrapcheck
	}

	return nil
}

// Close fecha a conexão com o banco de dados.
func Close(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err) //nolint:wrapcheck
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err) //nolint:wrapcheck
	}

	return nil
}
