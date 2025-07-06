package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"golang/internal/config"
	"golang/internal/middleware"
	"golang/internal/services"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// Server representa o servidor HTTP.
type Server struct {
	config      *config.Config
	db          *gorm.DB
	logger      *middleware.Logger
	router      *gin.Engine
	server      *http.Server
	tempService *services.TemperatureService
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
		config:      cfg,
		db:          db,
		logger:      logger,
		router:      router,
		tempService: services.NewTemperatureService(),
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

	// Rotas de temperatura
	temperature := v1.Group("/temperature")
	temperature.POST("/convert", s.convertTemperature)
	temperature.GET("/convert/:value/:from_unit", s.convertTemperatureGet)
	temperature.GET("/convert/:value/:from_unit/all", s.getAllConversions)

	// Adicione mais rotas aqui
	// users := v1.Group("/users")
	// users.GET("/", s.getUsers)
	// users.POST("/", s.createUser)
	// users.GET("/:id", s.getUser)
	// users.PUT("/:id", s.updateUser)
	// users.DELETE("/:id", s.deleteUser)
} //nolint:wsl

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

// convertTemperature converte temperatura via POST
func (s *Server) convertTemperature(c *gin.Context) {
	var req services.TemperatureConversionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := s.tempService.ConvertTemperature(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao converter temperatura",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// convertTemperatureGet converte temperatura via GET
func (s *Server) convertTemperatureGet(c *gin.Context) {
	valueStr := c.Param("value")
	fromUnit := c.Param("from_unit")
	toUnit := c.Query("to_unit")

	if toUnit == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parâmetro 'to_unit' é obrigatório",
		})
		return
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Valor inválido",
			"details": err.Error(),
		})
		return
	}

	req := services.TemperatureConversionRequest{
		Value:    value,
		FromUnit: fromUnit,
		ToUnit:   toUnit,
	}

	resp, err := s.tempService.ConvertTemperature(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao converter temperatura",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// getAllConversions retorna todas as conversões para um valor
func (s *Server) getAllConversions(c *gin.Context) {
	valueStr := c.Param("value")
	fromUnit := c.Param("from_unit")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Valor inválido",
			"details": err.Error(),
		})
		return
	}

	resp, err := s.tempService.GetAllConversions(value, fromUnit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao converter temperatura",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
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

	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err) //nolint:wrapcheck
	}

	return nil
}

// Shutdown desliga o servidor graciosamente.
func (s *Server) Shutdown(ctx context.Context) error {
	if s.server != nil {
		if err := s.server.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown server: %w", err) //nolint:wrapcheck
		}
	}

	return nil
}

// GetRouter retorna o router do servidor (usado para testes).
func (s *Server) GetRouter() *gin.Engine {
	return s.router
}
