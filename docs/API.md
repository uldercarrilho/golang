# API Documentation

## Visão Geral

Esta é a documentação da API REST do servidor Go. A API segue os padrões RESTful e retorna respostas em formato JSON.

## Base URL

```
http://localhost:8080
```

## Autenticação

Atualmente, a API não requer autenticação. Para endpoints protegidos no futuro, será implementado JWT.

## Endpoints

### Health Check

#### GET /health

Verifica o status de saúde da aplicação.

**Resposta:**
```json
{
  "status": "ok",
  "timestamp": "2024-01-01T12:00:00Z",
  "service": "golang-api"
}
```

**Status Codes:**
- `200 OK` - Servidor funcionando normalmente

### Hello World

#### GET /api/v1/hello

Endpoint de exemplo que retorna uma mensagem de saudação.

**Resposta:**
```json
{
  "message": "Hello, World!",
  "time": "2024-01-01T12:00:00Z"
}
```

**Status Codes:**
- `200 OK` - Requisição processada com sucesso

## Códigos de Status HTTP

- `200 OK` - Requisição processada com sucesso
- `201 Created` - Recurso criado com sucesso
- `400 Bad Request` - Dados inválidos na requisição
- `401 Unauthorized` - Autenticação necessária
- `403 Forbidden` - Acesso negado
- `404 Not Found` - Recurso não encontrado
- `500 Internal Server Error` - Erro interno do servidor

## Headers

### Requisição
```
Content-Type: application/json
Accept: application/json
```

### Resposta
```
Content-Type: application/json
```

## Exemplos de Uso

### Usando curl

```bash
# Health check
curl -X GET http://localhost:8080/health

# Hello endpoint
curl -X GET http://localhost:8080/api/v1/hello
```

### Usando JavaScript

```javascript
// Health check
fetch('http://localhost:8080/health')
  .then(response => response.json())
  .then(data => console.log(data));

// Hello endpoint
fetch('http://localhost:8080/api/v1/hello')
  .then(response => response.json())
  .then(data => console.log(data));
```

## Rate Limiting

Atualmente não há rate limiting implementado. Será adicionado em versões futuras.

## Versionamento

A API usa versionamento por URL. A versão atual é `v1`.

## Suporte

Para suporte técnico, entre em contato através dos canais oficiais do projeto. 