package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config representa as configurações da aplicação
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Log      LogConfig
}

// ServerConfig configurações do servidor
type ServerConfig struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

// DatabaseConfig configurações do banco de dados
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// LogConfig configurações de log
type LogConfig struct {
	Level string
}

// Load carrega as configurações do ambiente
func Load() (*Config, error) {
	// Carregar variáveis de ambiente do arquivo .env se existir
	_ = godotenv.Load()

	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  getEnvAsInt("READ_TIMEOUT", 30),
			WriteTimeout: getEnvAsInt("WRITE_TIMEOUT", 30),
			IdleTimeout:  getEnvAsInt("IDLE_TIMEOUT", 60),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "golang_app"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Log: LogConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
	}, nil
}

// getEnv obtém uma variável de ambiente ou retorna um valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt obtém uma variável de ambiente como inteiro ou retorna um valor padrão
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
