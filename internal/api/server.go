package api

import (
	"context"
	"net/http"
	"time"

	"golang/internal/config"
	"golang/internal/middleware"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// Server representa o servidor HTTP.
type Server struct {
	config *config.Config
	db     *gorm.DB
	logger *middleware.Logger
	router *gin.Engine
	server *http.Server
}

// NewServer cria uma nova instância do servidor.
func NewServer(cfg *config.Config, db *gorm.DB, logger *middleware.Logger) *Server {
	// Configurar modo do Gin
	if cfg.Log.Level == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Aplicar middlewares
	router.Use(middleware.RecoveryMiddleware(logger))
	router.Use(middleware.LoggingMiddleware(logger))
	router.Use(middleware.CORSMiddleware())

	server := &Server{
		config: cfg,
		db:     db,
		logger: logger,
		router: router,
	}

	// Configurar rotas
	server.setupRoutes()

	return server
}

// setupRoutes configura as rotas da aplicação.
func (s *Server) setupRoutes() {
	// Health check
	s.router.GET("/health", s.healthCheck)

	// API v1
	v1 := s.router.Group("/api/v1")
	// Exemplo de rota
	v1.GET("/hello", s.helloHandler)
	// Adicione mais rotas aqui
	// users := v1.Group("/users")
	// users.GET("/", s.getUsers)
	// users.POST("/", s.createUser)
	// users.GET("/:id", s.getUser)
	// users.PUT("/:id", s.updateUser)
	// users.DELETE("/:id", s.deleteUser)
}

// healthCheck retorna o status de saúde da aplicação.
func (s *Server) healthCheck(c *gin.Context) {
	status := gin.H{
		"status":    "ok",
		"timestamp": time.Now().UTC(),
		"service":   "golang-api",
	}

	// Adicionar status do banco se disponível
	if s.db != nil {
		status["database"] = "connected"
	} else {
		status["database"] = "disconnected"
	}

	c.JSON(http.StatusOK, status)
}

// helloHandler exemplo de handler.
func (s *Server) helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
		"time":    time.Now().UTC(),
	})
}

// Start inicia o servidor HTTP.
func (s *Server) Start(addr string) error {
	s.server = &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  time.Duration(s.config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(s.config.Server.IdleTimeout) * time.Second,
	}

	s.logger.Infof("Server starting on %s", addr)
	return s.server.ListenAndServe()
}

// Shutdown desliga o servidor graciosamente.
func (s *Server) Shutdown(ctx context.Context) error {
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

// GetRouter retorna o router do servidor (usado para testes).
func (s *Server) GetRouter() *gin.Engine {
	return s.router
}
