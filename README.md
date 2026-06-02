# Go Task Manager API

API REST para gerenciamento de tarefas (To-Do List), desenvolvida em **Go**. O projeto segue o [Standard Go Project Layout](https://github.com/golang-standards/project-layout) com arquitetura em camadas (handler → service → repository), ORM com GORM e configuração via variáveis de ambiente.

## Tecnologias

| Tecnologia | Uso |
| :--- | :--- |
| [Go](https://go.dev/) | Linguagem de Programação |
| [Gin](https://github.com/gin-gonic/gin) | Framework HTTP |
| [GORM](https://gorm.io/) | ORM (Mapeamento Objeto-Relacional) |
| [PostgreSQL](https://www.postgresql.org/) | Banco de dados |
| [godotenv](https://github.com/joho/godotenv) | Carregamento de `.env` em desenvolvimento |
| [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Hashing seguro de senhas |
| [JWT-Go (v5)](https://github.com/golang-jwt/jwt) | Emissão e validação de tokens de acesso |


## Estrutura do projeto

```
api_rest/
├── cmd/api/                 # Ponto de entrada da aplicação (main.go)
├── internal/
│   ├── config/              # Variáveis de ambiente e DSN
│   ├── database/            # Conexão e migrations automáticas
│   ├── middleware/          # Interceptador de segurança (AuthRequired JWT)
│   ├── model/               # Entidades e structs (Tarefa, Usuario)
│   ├── repository/          # Camada de persistência e banco de dados (GORM)
│   ├── service/             # Regras de negócio, Bcrypt e assinatura (JwtKey)
│   ├── handler/             # Handlers HTTP para processamento de JSON
│   └── router/              # Configuração e agrupamento de rotas Gin
├── .env.example             # Modelo de configuração local
├── .gitignore               # Arquivos ignorados no versionamento (.env)
└── go.mod                   # Gerenciador de módulos e dependências

```

## Pré-requisitos

- **Go** 1.26 ou superior
- **PostgreSQL** em execução (porta configurável via `DB_PORT`)

## Configuração

A conexão com o banco é definida por variáveis de ambiente. Em desenvolvimento, copie o arquivo de exemplo e ajuste os valores:

```bash
cp .env.example .env
```

| Variável | Descrição | Padrão |
| :--- | :--- | :--- |
| `DB_HOST` | Host do PostgreSQL | `localhost` |
| `DB_USER` | Usuário do banco | `postgres` |
| `DB_PASSWORD` | Senha do banco | `postgres` |
| `DB_NAME` | Nome do banco | `postgres` |
| `DB_PORT` | Porta do banco | `5432` |
| `DB_SSLMODE` | Modo SSL (`disable`, `require`, etc.) | `disable` |
| `JWT_SECRET` | Chave de assinatura dos tokens (`opcional se houver fallback`) | `variável de sistema` |

O arquivo `.env` é ignorado pelo Git. Em produção, defina as variáveis diretamente no ambiente (Docker, Kubernetes, CI, etc.) — o `.env` não é obrigatório.

Se o `.env` não existir, a API usa as variáveis do sistema operacional ou os valores padrão da tabela acima.

## Como executar

1. Clone o repositório:

   ```bash
   git clone https://github.com/TrentinJr/API_Rest.git
   cd API_Rest
   ```

2. Configure o banco (veja [Configuração](#configuração)).

3. Instale as dependências e execute:

   ```bash
   go mod download
   go run ./cmd/api
   ```

A API sobe em `http://localhost:8080`. As migrations das tabelas são aplicadas automaticamente na inicialização.

## Endpoints

Base URL: `http://localhost:8080`

👤 Autenticação e Usuários (Rotas Públicas)
| Método | Endpoint       | Descrição                                             | Corpo da Requisição (JSON)        |
| `POST` |`/usuario`      | Cadastra um novo usuário com senha criptografada      | `{"email": "...", "senha": "..."}`|
| `POST` |`/usuario/login`| Autentica o usuário e gera um Token JWT válido por 24h|``{"email": "...", "senha": "..."}`|

📝 Gerenciamento de Tarefas (Rotas Protegidas por JWT)
⚠️ Obrigatório: Todas as requisições abaixo necessitam do cabeçalho Authorization: Bearer <SEU_TOKEN_JWT>.

| Método | Endpoint             | Descrição                                      | Corpo da Requisição (`JSON`)                               |
|`GET`   |`/tarefas`            |`Lista todas as tarefas do banco de dados`      |`nenhum`                                                    |
|`POST`  |`/tarefas`            |`Cria uma nova tarefa no sistema`               |`{"titulo": "...", "descricao": "...", "status": "..."}`    |
|`PUT`   |`/tarefas/:id`        |`Atualiza título, descrição ou dados completos` |`{"titulo": "...", "descricao": "...", "status": "..."}`    |
|`PATCH` |`/tarefas/:id/status` |`Atualiza exclusivamente o status da tarefa`    |`{"status": "concluido"}`                                   |
|`DELETE`|`/tarefas/:id`        |`Remove uma tarefa de forma definitiva`         |`nenhum`                                                    |

### Exemplos de Requisição por cURL

1. Realizar Login e obter Token:
curl -X POST http://localhost:8080/usuarios/login \
  -H "Content-Type: application/json" \
  -d '{"email": "teste@email.com", "senha": "senha123"}'


2. Listar tarefas (Exemplo protegida):
curl http://localhost:8080/tarefas \
  -H "Authorization: Bearer SUBST_PELO_SEU_TOKEN_JWT"


3. Alterar status da tarefa:
curl -X PATCH http://localhost:8080/tarefas/1/status \
  -H "Authorization: Bearer SUBST_PELO_SEU_TOKEN_JWT" \
  -H "Content-Type: application/json" \
  -d '{"status": "concluido"}'


### Respostas de erro comuns

| Status | Situação |
| :--- | :--- |
| `400` | JSON inválido ou ID malformado |
| `401` | Unauthorized |
| `404` | Tarefa não encontrada |
| `500` | Erro interno (ex.: falha no banco) |

## Testando a API

Você pode testar os endpoints com [Postman](https://www.postman.com/), [Insomnia](https://insomnia.rest/) ou os exemplos `curl` acima.

## Licença

Este projeto é de uso livre para fins de estudo e desenvolvimento.
