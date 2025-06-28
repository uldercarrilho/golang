# Changelog

Todas as mudanças notáveis neste projeto serão documentadas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/),
e este projeto adere ao [Versionamento Semântico](https://semver.org/lang/pt-BR/).

## [1.24.4] - 2024-12-28

### Adicionado
- Suporte ao Go 1.24.4
- Método `GetRouter()` na struct `Server` para facilitar testes
- Melhorias na estrutura de testes de integração

### Atualizado
- Go: 1.23 → 1.24.4
- Gin: v1.9.1 → v1.10.1
- GORM: v1.25.4 → v1.30.0
- GORM Driver PostgreSQL: v1.5.2 → v1.6.0
- Testify: v1.8.4 → v1.10.0
- Godotenv: v1.4.0 → v1.5.1
- Todas as dependências indiretas atualizadas para as versões mais recentes

### Corrigido
- Problemas de importação do GORM nos testes de integração
- Acesso ao campo privado `router` nos testes
- Método de fechamento de conexão do banco de dados nos testes

### Testado
- ✅ Todos os testes unitários passando
- ✅ Testes de integração funcionando corretamente
- ✅ Validação de emails, senhas, UUIDs, telefones e CPFs
- ✅ Endpoints de health check e hello world
- ✅ Conexão com banco de dados PostgreSQL

## [1.23.0] - 2024-01-01

### Adicionado
- Estrutura inicial do projeto
- Configuração básica do servidor HTTP
- Sistema de logging estruturado
- Conexão com banco de dados PostgreSQL
- Testes automatizados
- Documentação da API
- Containerização com Docker

### Tecnologias
- Go 1.23
- Gin v1.9.1
- GORM v1.25.4
- PostgreSQL
- Logrus
- Testify v1.8.4 