package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
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

// TestConvertTemperature testa o endpoint de conversão de temperatura via POST
func TestConvertTemperature(t *testing.T) {
	cfg := &config.Config{
		Server:   config.ServerConfig{},
		Database: config.DatabaseConfig{},
		Log: config.LogConfig{
			Level: "info",
		},
	}

	logger := middleware.NewLogger()
	var db *gorm.DB
	server := NewServer(cfg, db, logger)

	// Teste de conversão Celsius para Fahrenheit
	jsonBody := `{"value": 25.0, "from_unit": "celsius", "to_unit": "fahrenheit"}`
	req, err := http.NewRequestWithContext(context.Background(), "POST", "/api/v1/temperature/convert", strings.NewReader(jsonBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	server.GetRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "77")
	assert.Contains(t, w.Body.String(), "fahrenheit")
}

// TestConvertTemperatureGet testa o endpoint de conversão de temperatura via GET
func TestConvertTemperatureGet(t *testing.T) {
	cfg := &config.Config{
		Server:   config.ServerConfig{},
		Database: config.DatabaseConfig{},
		Log: config.LogConfig{
			Level: "info",
		},
	}

	logger := middleware.NewLogger()
	var db *gorm.DB
	server := NewServer(cfg, db, logger)

	// Teste de conversão Celsius para Fahrenheit via GET
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/api/v1/temperature/convert/25/celsius?to_unit=fahrenheit", http.NoBody)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	server.GetRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "77")
	assert.Contains(t, w.Body.String(), "fahrenheit")
}

// TestGetAllConversions testa o endpoint de todas as conversões
func TestGetAllConversions(t *testing.T) {
	cfg := &config.Config{
		Server:   config.ServerConfig{},
		Database: config.DatabaseConfig{},
		Log: config.LogConfig{
			Level: "info",
		},
	}

	logger := middleware.NewLogger()
	var db *gorm.DB
	server := NewServer(cfg, db, logger)

	// Teste de todas as conversões para Celsius
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/api/v1/temperature/convert/25/celsius/all", http.NoBody)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	server.GetRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "25")
	assert.Contains(t, w.Body.String(), "77")
	assert.Contains(t, w.Body.String(), "298.15")
}

// TestConvertTemperatureInvalidData testa validação de dados inválidos
func TestConvertTemperatureInvalidData(t *testing.T) {
	cfg := &config.Config{
		Server:   config.ServerConfig{},
		Database: config.DatabaseConfig{},
		Log: config.LogConfig{
			Level: "info",
		},
	}

	logger := middleware.NewLogger()
	var db *gorm.DB
	server := NewServer(cfg, db, logger)

	// Teste com dados inválidos
	jsonBody := `{"value": "invalid", "from_unit": "celsius", "to_unit": "fahrenheit"}`
	req, err := http.NewRequestWithContext(context.Background(), "POST", "/api/v1/temperature/convert", strings.NewReader(jsonBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	server.GetRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}
