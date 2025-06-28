# Golang Server - Documentação

## Visão Geral

Este é um servidor de aplicação em Go que segue as melhores práticas da comunidade. O projeto utiliza uma arquitetura limpa e modular, com separação clara de responsabilidades.

## Tecnologias Utilizadas

- **Go 1.21** - Linguagem principal
- **Gin** - Framework web
- **GORM** - ORM para banco de dados
- **PostgreSQL** - Banco de dados principal
- **Logrus** - Sistema de logging
- **Testify** - Framework de testes
- **Docker** - Containerização

## Estrutura do Projeto

```
golang/
├── cmd/                    # Entry points da aplicação
│   └── server/            # Servidor principal
├── internal/              # Código privado da aplicação
│   ├── api/              # Handlers HTTP
│   ├── config/           # Configurações
│   ├── database/         # Camada de dados
│   ├── middleware/       # Middlewares HTTP
│   ├── models/           # Modelos de dados
│   └── services/         # Lógica de negócio
├── pkg/                  # Código público reutilizável
├── test/                 # Testes de integração
├── docs/                 # Documentação
├── .github/              # GitHub Actions
├── go.mod                # Dependências
├── Makefile              # Comandos de build
├── Dockerfile            # Containerização
└── docker-compose.yml    # Orquestração
```

## Pré-requisitos

- Go 1.21 ou superior
- PostgreSQL 15 ou superior
- Docker e Docker Compose (opcional)

## Instalação e Configuração

### 1. Clone o repositório

```bash
git clone <repository-url>
cd golang
```

### 2. Instale as dependências

```bash
go mod tidy
```

### 3. Configure as variáveis de ambiente

```bash
cp env.example .env
# Edite o arquivo .env com suas configurações
```

### 4. Configure o banco de dados

```bash
# Usando Docker Compose (recomendado)
docker-compose up -d postgres

# Ou configure um PostgreSQL local
```

### 5. Execute a aplicação

```bash
# Desenvolvimento
make run

# Ou usando go diretamente
go run cmd/server/main.go
```

## Comandos Úteis

### Makefile

```bash
# Build da aplicação
make build

# Executar testes
make test

# Executar testes com coverage
make test-coverage

# Linting
make lint

# Formatação de código
make format

# Limpar arquivos de build
make clean

# Instalar dependências
make deps

# Ver todos os comandos
make help
```

### Docker

```bash
# Build da imagem
make docker-build

# Executar container
make docker-run

# Ou usando docker-compose
docker-compose up -d
```

## Desenvolvimento

### Adicionando Novos Endpoints

1. Crie o modelo em `internal/models/`
2. Crie o serviço em `internal/services/`
3. Adicione o handler em `internal/api/server.go`
4. Escreva testes em `internal/api/server_test.go`

### Exemplo de Novo Endpoint

```go
// Em internal/api/server.go
func (s *Server) setupRoutes() {
    // ... código existente ...
    
    users := v1.Group("/users")
    {
        users.GET("/", s.getUsers)
        users.POST("/", s.createUser)
        users.GET("/:id", s.getUser)
        users.PUT("/:id", s.updateUser)
        users.DELETE("/:id", s.deleteUser)
    }
}
```

### Testes

```bash
# Executar todos os testes
go test ./...

# Executar testes com coverage
go test -cover ./...

# Executar testes de um pacote específico
go test ./internal/api
```

## Deploy

### Docker

```bash
# Build da imagem
docker build -t golang-server .

# Executar container
docker run -p 8080:8080 golang-server
```

### Docker Compose

```bash
# Executar todos os serviços
docker-compose up -d

# Ver logs
docker-compose logs -f app
```

## Monitoramento

### Health Check

```bash
curl http://localhost:8080/health
```

### Logs

Os logs são estruturados em JSON e incluem:
- Timestamp
- Nível de log
- Mensagem
- Campos adicionais (IP, user-agent, etc.)

## Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a MIT License - veja o arquivo [LICENSE](LICENSE) para detalhes.

## Suporte

Para suporte técnico ou dúvidas, abra uma issue no repositório do projeto. 