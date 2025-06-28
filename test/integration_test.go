package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang/internal/api"
	"golang/internal/config"
	"golang/internal/database"
	"golang/internal/middleware"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestServer representa um servidor de teste
type TestServer struct {
	Server *api.Server
	DB     *gorm.DB
}

// NewTestServer cria um novo servidor para testes
func NewTestServer(t *testing.T) *TestServer {
	// Configuração de teste
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port:         "8080",
			ReadTimeout:  30,
			WriteTimeout: 30,
			IdleTimeout:  60,
		},
		Database: config.DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "password",
			DBName:   "golang_test",
			SSLMode:  "disable",
		},
		Log: config.LogConfig{
			Level: "error", // Reduzir logs durante testes
		},
	}

	// Conectar ao banco de dados de teste
	db, err := database.Connect(cfg.Database)
	require.NoError(t, err)

	// Criar logger
	logger := middleware.NewLogger()

	// Criar servidor
	server := api.NewServer(cfg, db, logger)

	return &TestServer{
		Server: server,
		DB:     db,
	}
}

// TestHealthCheck testa o endpoint de health check
func TestHealthCheck(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.DB.Close()

	// Criar requisição
	req, err := http.NewRequest("GET", "/health", nil)
	require.NoError(t, err)

	// Criar response recorder
	w := httptest.NewRecorder()

	// Executar requisição
	ts.Server.router.ServeHTTP(w, req)

	// Verificar resposta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

// TestHelloEndpoint testa o endpoint hello
func TestHelloEndpoint(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.DB.Close()

	// Criar requisição
	req, err := http.NewRequest("GET", "/api/v1/hello", nil)
	require.NoError(t, err)

	// Criar response recorder
	w := httptest.NewRecorder()

	// Executar requisição
	ts.Server.router.ServeHTTP(w, req)

	// Verificar resposta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Hello, World!")
}
