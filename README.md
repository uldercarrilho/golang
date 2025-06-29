# Golang Server

[![Go Version](https://img.shields.io/badge/Go-1.24.4-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/golang)](https://goreportcard.com/report/github.com/yourusername/golang)

Um servidor de aplicação em Go que segue as melhores práticas da comunidade, com arquitetura limpa e modular.

## 🚀 Características

- **Arquitetura Limpa**: Separação clara de responsabilidades
- **Framework Moderno**: Gin v1.10.1 para HTTP routing
- **ORM Robusto**: GORM v1.30.0 para operações de banco de dados
- **Logging Estruturado**: Logrus para logs em JSON
- **Testes Automatizados**: Cobertura completa com Testify v1.10.0
- **Containerização**: Docker multi-stage para produção
- **CI/CD**: GitHub Actions para automação
- **Documentação**: API docs e guias de desenvolvimento

## 📋 Pré-requisitos

- Go 1.24.4 ou superior
- PostgreSQL 15 ou superior
- Docker e Docker Compose (opcional)

## 🛠️ Instalação Rápida

### Usando Docker Compose (Recomendado)

```bash
# Clone o repositório
git clone <repository-url>
cd golang

# Execute com Docker Compose
docker-compose up -d

# Acesse a aplicação
curl http://localhost:8080/health
```

### Instalação Local

```bash
# Clone o repositório
git clone <repository-url>
cd golang

# Instale dependências
go mod tidy

# Configure variáveis de ambiente
cp env.example .env

# Execute a aplicação
make run
```

## 📖 Documentação

- [Guia de Desenvolvimento](docs/README.md)
- [Documentação da API](docs/API.md)
- [Estrutura do Projeto](docs/README.md#estrutura-do-projeto)
- [Changelog](CHANGELOG.md)

## 🔄 Atualizações Recentes

### Versão 1.24.4 (Dezembro 2024)
- ✅ **Go 1.24.4** - Versão mais recente do Go
- ✅ **Gin v1.10.1** - Framework web atualizado
- ✅ **GORM v1.30.0** - ORM com melhorias de performance
- ✅ **Testify v1.10.0** - Framework de testes atualizado
- ✅ **Todas as dependências** atualizadas para as versões mais recentes
- ✅ **Testes automatizados** funcionando perfeitamente

Veja o [CHANGELOG.md](CHANGELOG.md) para detalhes completos das mudanças.

## 🏗️ Estrutura do Projeto

```
golang/
├── cmd/server/           # Entry point principal
├── internal/             # Código privado da aplicação
│   ├── api/             # Handlers HTTP
│   ├── config/          # Configurações
│   ├── database/        # Camada de dados
│   ├── middleware/      # Middlewares HTTP
│   ├── models/          # Modelos de dados
│   └── services/        # Lógica de negócio
├── test/                # Testes de integração
├── docs/                # Documentação
├── .github/workflows/   # CI/CD
├── Dockerfile           # Containerização
└── docker-compose.yml   # Orquestração
```

## 🧪 Testes

```bash
# Executar todos os testes
make test

# Executar testes com coverage
make test-coverage

# Executar testes de integração
go test ./test/...
```

## 🐳 Docker

```bash
# Build da imagem
make docker-build

# Executar container
make docker-run

# Usando docker-compose
docker-compose up -d
```

## 🚀 CI/CD

Este projeto utiliza GitHub Actions para automação de CI/CD. O pipeline inclui:

- ✅ **Testes Automatizados** - Executa todos os testes em cada push
- ✅ **Linting** - Verifica qualidade do código
- ✅ **Build** - Compila a aplicação
- ✅ **Docker** - Constrói e publica imagem no Docker Hub

### Configuração do Docker Hub

Para que o pipeline funcione corretamente, configure os secrets no GitHub:

1. Vá para **Settings** > **Secrets and variables** > **Actions**
2. Adicione os secrets:
   - `DOCKER_USERNAME` - Seu usuário do Docker Hub
   - `DOCKER_PASSWORD` - Sua senha ou token de acesso

> 📖 Veja o [Guia de Deploy](docs/DEPLOYMENT.md) para instruções detalhadas.

### Status do Pipeline

O pipeline é executado automaticamente em:
- Push para `main` e `develop`
- Pull requests para `main`

## 📊 Monitoramento

### Health Check

```bash
curl http://localhost:8080/health
```

### Endpoints Disponíveis

- `GET /health` - Status de saúde da aplicação
- `GET /api/v1/hello` - Endpoint de exemplo

## 🔧 Desenvolvimento

### Comandos Úteis

```bash
# Ver todos os comandos disponíveis
make help

# Formatação de código
make format

# Linting
make lint

# Instalar ferramentas de desenvolvimento
make install-tools
```

### Adicionando Novos Endpoints

1. Crie o modelo em `internal/models/`
2. Crie o serviço em `internal/services/`
3. Adicione o handler em `internal/api/server.go`
4. Escreva testes em `internal/api/server_test.go`

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está licenciado sob a MIT License - veja o arquivo [LICENSE](LICENSE) para detalhes.

## 🆘 Suporte

- 📖 [Documentação](docs/README.md)
- 🐛 [Issues](https://github.com/yourusername/golang/issues)
- 💬 [Discussions](https://github.com/yourusername/golang/discussions)

## 🙏 Agradecimentos

- [Gin](https://github.com/gin-gonic/gin) - Framework web
- [GORM](https://gorm.io/) - ORM para Go
- [Logrus](https://github.com/sirupsen/logrus) - Sistema de logging
- [Testify](https://github.com/stretchr/testify) - Framework de testes
