# Variáveis
BINARY_NAME=golang-server
BUILD_DIR=build
MAIN_PATH=./cmd/server

# Comandos principais
.PHONY: build run test clean lint format help

# Build da aplicação
build:
	@echo "Building application..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

# Executar a aplicação
run:
	@echo "Running application..."
	go run $(MAIN_PATH)

# Executar testes
test:
	@echo "Running tests..."
	go test -v ./...

# Executar testes com coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Limpar arquivos de build
clean:
	@echo "Cleaning build files..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html

# Linting com golangci-lint
lint:
	@echo "Running linter..."
	golangci-lint run

# Formatação de código
format:
	@echo "Formatting code..."
	go fmt ./...
	go vet ./...

# Instalar dependências
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# Gerar documentação
docs:
	@echo "Generating documentation..."
	godoc -http=:6060

# Docker build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME) .

# Docker run
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(BINARY_NAME)

# Instalar ferramentas de desenvolvimento
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/godoc@latest

# Ajuda
help:
	@echo "Available commands:"
	@echo "  build         - Build the application"
	@echo "  run           - Run the application"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  clean         - Clean build files"
	@echo "  lint          - Run linter"
	@echo "  format        - Format code"
	@echo "  deps          - Install dependencies"
	@echo "  docs          - Generate documentation"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-run    - Run Docker container"
	@echo "  install-tools - Install development tools"
	@echo "  help          - Show this help" 