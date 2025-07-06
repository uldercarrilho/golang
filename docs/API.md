# API Documentation

## Visão Geral

Esta é a documentação da API REST do servidor Go. A API segue os padrões RESTful e retorna respostas em formato JSON.

## Atualizações Recentes

### Versão 1.24.4 (Dezembro 2024)
- ✅ Atualizado para Go 1.24.4
- ✅ Gin atualizado para v1.10.1
- ✅ GORM atualizado para v1.30.0
- ✅ Testify atualizado para v1.10.0
- ✅ Todas as dependências atualizadas para as versões mais recentes
- ✅ Testes automatizados funcionando corretamente
- ✅ Melhorias na estrutura de testes de integração
- ✅ **NOVO**: Endpoint de conversão de temperatura (Kelvin, Celsius, Fahrenheit)
- ✅ **NOVO**: Testes automatizados para conversão de temperatura

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

### Conversão de Temperatura

#### POST /api/v1/temperature/convert

Converte uma temperatura de uma unidade para outra via POST.

**Corpo da Requisição:**
```json
{
  "value": 25.0,
  "from_unit": "celsius",
  "to_unit": "fahrenheit"
}
```

**Unidades Suportadas:**
- `celsius` - Graus Celsius (°C)
- `fahrenheit` - Graus Fahrenheit (°F)
- `kelvin` - Kelvin (K)

**Resposta:**
```json
{
  "original_value": 25.0,
  "original_unit": "celsius",
  "converted_value": 77.0,
  "converted_unit": "fahrenheit",
  "formula": "°F = (°C × 9/5) + 32 = (25.00 × 9/5) + 32 = 77.00"
}
```

**Status Codes:**
- `200 OK` - Conversão realizada com sucesso
- `400 Bad Request` - Dados inválidos na requisição

#### GET /api/v1/temperature/convert/:value/:from_unit

Converte uma temperatura de uma unidade para outra via GET.

**Parâmetros:**
- `value` - Valor da temperatura (número)
- `from_unit` - Unidade de origem (celsius, fahrenheit, kelvin)
- `to_unit` - Unidade de destino (query parameter)

**Exemplo:**
```
GET /api/v1/temperature/convert/25/celsius?to_unit=fahrenheit
```

**Resposta:**
```json
{
  "original_value": 25.0,
  "original_unit": "celsius",
  "converted_value": 77.0,
  "converted_unit": "fahrenheit",
  "formula": "°F = (°C × 9/5) + 32 = (25.00 × 9/5) + 32 = 77.00"
}
```

#### GET /api/v1/temperature/convert/:value/:from_unit/all

Retorna todas as conversões possíveis para um valor dado.

**Parâmetros:**
- `value` - Valor da temperatura (número)
- `from_unit` - Unidade de origem (celsius, fahrenheit, kelvin)

**Exemplo:**
```
GET /api/v1/temperature/convert/25/celsius/all
```

**Resposta:**
```json
{
  "original_value": 25.0,
  "original_unit": "celsius",
  "conversions": {
    "celsius": 25.0,
    "fahrenheit": 77.0,
    "kelvin": 298.15
  },
  "formulas": {
    "celsius": "Mesma unidade, sem conversão necessária",
    "fahrenheit": "°F = (°C × 9/5) + 32 = (25.00 × 9/5) + 32 = 77.00",
    "kelvin": "K = °C + 273.15 = 25.00 + 273.15 = 298.15"
  }
}
```

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

# Conversão de temperatura via POST
curl -X POST http://localhost:8080/api/v1/temperature/convert \
  -H "Content-Type: application/json" \
  -d '{"value": 25.0, "from_unit": "celsius", "to_unit": "fahrenheit"}'

# Conversão de temperatura via GET
curl -X GET "http://localhost:8080/api/v1/temperature/convert/25/celsius?to_unit=fahrenheit"

# Todas as conversões para um valor
curl -X GET http://localhost:8080/api/v1/temperature/convert/25/celsius/all
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

// Conversão de temperatura via POST
fetch('http://localhost:8080/api/v1/temperature/convert', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    value: 25.0,
    from_unit: 'celsius',
    to_unit: 'fahrenheit'
  })
})
.then(response => response.json())
.then(data => console.log(data));

// Conversão de temperatura via GET
fetch('http://localhost:8080/api/v1/temperature/convert/25/celsius?to_unit=fahrenheit')
  .then(response => response.json())
  .then(data => console.log(data));

// Todas as conversões para um valor
fetch('http://localhost:8080/api/v1/temperature/convert/25/celsius/all')
  .then(response => response.json())
  .then(data => console.log(data));
```

## Rate Limiting

Atualmente não há rate limiting implementado. Será adicionado em versões futuras.

## Versionamento

A API usa versionamento por URL. A versão atual é `v1`.

## Suporte

Para suporte técnico, entre em contato através dos canais oficiais do projeto. 