# Guia de Deploy e CI/CD

## Configuração do Docker Hub para CI/CD

### 1. Criar conta no Docker Hub

1. Acesse [hub.docker.com](https://hub.docker.com)
2. Crie uma conta gratuita
3. Faça login na sua conta

### 2. Configurar Secrets no GitHub

Para que o pipeline CI/CD funcione corretamente, você precisa configurar os seguintes secrets no seu repositório GitHub:

1. Acesse seu repositório no GitHub
2. Vá para **Settings** > **Secrets and variables** > **Actions**
3. Clique em **New repository secret**
4. Adicione os seguintes secrets:

#### DOCKER_USERNAME
- **Name**: `DOCKER_USERNAME`
- **Value**: Seu nome de usuário do Docker Hub

#### DOCKER_PASSWORD
- **Name**: `DOCKER_PASSWORD`
- **Value**: Sua senha do Docker Hub ou um token de acesso

> **Nota**: Para maior segurança, recomenda-se usar um token de acesso ao invés da senha:
> 1. No Docker Hub, vá para **Account Settings** > **Security**
> 2. Clique em **New Access Token**
> 3. Dê um nome ao token (ex: "GitHub Actions")
> 4. Copie o token gerado e use como valor do secret `DOCKER_PASSWORD`

### 3. Verificar Configuração

Após configurar os secrets, o pipeline irá:

1. Executar os testes
2. Fazer build da aplicação
3. Fazer login no Docker Hub
4. Construir e fazer push da imagem Docker

### 4. Estrutura das Imagens

As imagens serão publicadas com as seguintes tags:
- `seu-usuario/golang-server:latest` - Versão mais recente
- `seu-usuario/golang-server:sha-commit` - Versão específica do commit

### 5. Troubleshooting

#### Erro: "Username and password required"
- Verifique se os secrets `DOCKER_USERNAME` e `DOCKER_PASSWORD` estão configurados
- Certifique-se de que as credenciais estão corretas
- Se estiver usando token, verifique se ele não expirou

#### Erro: "Authentication failed"
- Verifique se o usuário e senha/token estão corretos
- Certifique-se de que a conta do Docker Hub está ativa
- Verifique se não há restrições de IP na conta

### 6. Deploy Local

Para testar localmente:

```bash
# Build da aplicação
make build

# Build da imagem Docker
make docker-build

# Executar com Docker Compose
docker-compose up
```

### 7. Variáveis de Ambiente

Certifique-se de configurar as seguintes variáveis de ambiente para produção:

```bash
# Banco de dados
DB_HOST=your-db-host
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=your-db-name
DB_SSLMODE=require

# Servidor
PORT=8080
LOG_LEVEL=info

# Segurança
JWT_SECRET=your-jwt-secret
API_KEY=your-api-key
```

### 8. Monitoramento

A aplicação inclui um endpoint de health check:
- `GET /health` - Verifica se a aplicação está funcionando

### 9. Logs

Os logs são estruturados em JSON e incluem:
- Timestamp
- Nível de log
- Mensagem
- Método HTTP
- Path
- Status code
- Latência
- User agent
- IP do cliente 