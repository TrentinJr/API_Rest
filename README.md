# Go Task Manager API

API REST para gerenciamento de tarefas (To-Do List), desenvolvida em **Go**. O projeto segue o [Standard Go Project Layout](https://github.com/golang-standards/project-layout) com arquitetura em camadas (handler → service → repository), ORM com GORM e configuração via variáveis de ambiente.

## Tecnologias

| Tecnologia | Uso |
| :--- | :--- |
| [Go](https://go.dev/) | Linguagem |
| [Gin](https://github.com/gin-gonic/gin) | Framework HTTP |
| [GORM](https://gorm.io/) | ORM |
| [PostgreSQL](https://www.postgresql.org/) | Banco de dados |
| [godotenv](https://github.com/joho/godotenv) | Carregamento de `.env` em desenvolvimento |

## Estrutura do projeto

```
api_rest/
├── cmd/api/                 # Ponto de entrada da aplicação
├── internal/
│   ├── config/              # Variáveis de ambiente e DSN
│   ├── database/            # Conexão e migrations
│   ├── model/               # Entidades
│   ├── repository/          # Persistência (GORM)
│   ├── service/             # Regras de negócio
│   ├── handler/             # Handlers HTTP
│   └── router/              # Rotas Gin
├── .env.example             # Modelo de configuração local
└── go.mod
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

O arquivo `.env` é ignorado pelo Git. Em produção, defina as variáveis diretamente no ambiente (Docker, Kubernetes, CI, etc.) — o `.env` não é obrigatório.

Se o `.env` não existir, a API usa as variáveis do sistema operacional ou os valores padrão da tabela acima.

## Como executar

1. Clone o repositório:

   ```bash
   git clone https://github.com/azevedoguigo/API_Rest.git
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

| Método | Endpoint | Descrição |
| :--- | :--- | :--- |
| `POST` | `/tarefas` | Cria uma tarefa |
| `GET` | `/tarefas` | Lista todas as tarefas |
| `PUT` | `/tarefas/:id` | Atualiza título e descrição |
| `PATCH` | `/tarefas/:id/status` | Atualiza o status de conclusão |
| `DELETE` | `/tarefas/:id` | Remove uma tarefa |

### Exemplos de requisição

**Criar tarefa**

```bash
curl -X POST http://localhost:8080/tarefas \
  -H "Content-Type: application/json" \
  -d '{"titulo": "Estudar Go", "descricao": "Praticar GORM"}'
```

**Listar tarefas**

```bash
curl http://localhost:8080/tarefas
```

**Editar tarefa**

```bash
curl -X PUT http://localhost:8080/tarefas/1 \
  -H "Content-Type: application/json" \
  -d '{"titulo": "Novo título", "descricao": "Nova descrição"}'
```

**Alterar status**

```bash
curl -X PATCH http://localhost:8080/tarefas/1/status \
  -H "Content-Type: application/json" \
  -d '{"concluida": true}'
```

**Deletar tarefa**

```bash
curl -X DELETE http://localhost:8080/tarefas/1
```

### Respostas de erro comuns

| Status | Situação |
| :--- | :--- |
| `400` | JSON inválido ou ID malformado |
| `404` | Tarefa não encontrada |
| `500` | Erro interno (ex.: falha no banco) |

## Testando a API

Você pode testar os endpoints com [Postman](https://www.postman.com/), [Insomnia](https://insomnia.rest/) ou os exemplos `curl` acima.

## Licença

Este projeto é de uso livre para fins de estudo e desenvolvimento.
