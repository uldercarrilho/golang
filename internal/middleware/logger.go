package middleware

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger é um wrapper para o logrus.
type Logger struct {
	*logrus.Logger
}

// NewLogger cria uma nova instância do logger.
func NewLogger() *Logger {
	log := logrus.New()

	// Configurar formato de saída
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// Configurar nível de log baseado na variável de ambiente
	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	// Configurar saída
	log.SetOutput(os.Stdout)

	return &Logger{log}
}

// WithField adiciona um campo ao logger.
func (l *Logger) WithField(key string, value interface{}) *logrus.Entry {
	return l.Logger.WithField(key, value)
}

// WithFields adiciona múltiplos campos ao logger.
func (l *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.Logger.WithFields(fields)
}

// Writer retorna o writer do logger para uso com gin.RecoveryWithWriter.
func (l *Logger) Writer() *logrus.Logger {
	return l.Logger
}
