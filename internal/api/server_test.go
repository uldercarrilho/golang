package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang/internal/config"
	"golang/internal/middleware"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestNewServer testa a criação de um novo servidor
func TestNewServer(t *testing.T) {
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port:         "8080",
			ReadTimeout:  30,
			WriteTimeout: 30,
			IdleTimeout:  60,
		},
		Database: config.DatabaseConfig{},
		Log: config.LogConfig{
			Level: "info",
		},
	}

	logger := middleware.NewLogger()
	var db *gorm.DB // Mock do banco

	server := NewServer(cfg, db, logger)

	assert.NotNil(t, server)
	assert.NotNil(t, server.router)
	assert.Equal(t, cfg, server.config)
}

// TestHealthCheck testa o endpoint de health check
func TestHealthCheck(t *testing.T) {
	cfg := &config.Config{
		Server:   config.ServerConfig{},
		Database: config.DatabaseConfig{},
		Log: config.LogConfig{
			Level: "info",
		},
	}

	logger := middleware.NewLogger()
	var db *gorm.DB // Mock do banco

	server := NewServer(cfg, db, logger)

	// Criar requisição
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/health", http.NoBody)
	require.NoError(t, err)

	// Criar response recorder
	w := httptest.NewRecorder()

	// Executar requisição
	server.GetRouter().ServeHTTP(w, req)

	// Verificar resposta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
	assert.Contains(t, w.Body.String(), "golang-api")
}

// TestHelloHandler testa o endpoint hello
func TestHelloHandler(t *testing.T) {
	cfg := &config.Config{
		Server:   config.ServerConfig{},
		Database: config.DatabaseConfig{},
		Log: config.LogConfig{
			Level: "info",
		},
	}

	logger := middleware.NewLogger()
	var db *gorm.DB // Mock do banco

	server := NewServer(cfg, db, logger)

	// Criar requisição
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/api/v1/hello", http.NoBody)
	require.NoError(t, err)

	// Criar response recorder
	w := httptest.NewRecorder()

	// Executar requisição
	server.GetRouter().ServeHTTP(w, req)

	// Verificar resposta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Hello, World!")
}
