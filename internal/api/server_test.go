package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang/internal/config"
	"golang/internal/middleware"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// mockDB é um mock do banco de dados para testes
type mockDB struct{}

func (m *mockDB) AutoMigrate(dst ...interface{}) error { return nil }
func (m *mockDB) Close() error                         { return nil }

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
	req, err := http.NewRequest("GET", "/health", nil)
	require.NoError(t, err)

	// Criar response recorder
	w := httptest.NewRecorder()

	// Executar requisição
	server.router.ServeHTTP(w, req)

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
	req, err := http.NewRequest("GET", "/api/v1/hello", nil)
	require.NoError(t, err)

	// Criar response recorder
	w := httptest.NewRecorder()

	// Executar requisição
	server.router.ServeHTTP(w, req)

	// Verificar resposta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Hello, World!")
}
