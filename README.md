# Go Task Manager - Full Stack Application

Este é um ecossistema completo para gerenciamento de tarefas (To-Do List). O projeto é composto por uma **API REST robusta** desenvolvida em **Go (Golang)** conectada a um banco de dados PostgreSQL, integrada a uma **interface moderna e interativa** 

---

## Tecnologias

### Backend
| Tecnologia | Uso |
| :--- | :--- |
| [Go](https://go.dev/) | Linguagem de Programação |
| [Gin](https://github.com/gin-gonic/gin) | Framework HTTP |
| [GORM](https://gorm.io/) | ORM (Mapeamento Objeto-Relacional) |
| [PostgreSQL](https://www.postgresql.org/) | Banco de dados |
| [godotenv](https://github.com/joho/godotenv) | Carregamento de `.env` em desenvolvimento |
| [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Hashing seguro de senhas |
| [JWT-Go (v5)](https://github.com/golang-jwt/jwt) | Emissão e validação de tokens de acesso |

### Frontend
| Tecnologia     | Uso |
|:---            | :--- |
| [React](https://react.dev/)| Biblioteca para construção de interfaces |
| [Vite](https://vite.dev/)  | Bundler e servidor de desenvolvimento ultrarrápido |
| [TanStack Start](https://tanstack.com/router/v1/docs/start/overview)| Framework para controle de rotas e estado |
| [Tailwind CSS](https://tailwindcss.com/)| Estilização utilitária baseada em classes |



## 📁 Estrutura do Projeto

O projeto adota uma arquitetura em camadas no backend seguindo o *Standard Go Project Layout*, e centraliza todo o código visual na subpasta `frontend/`.

```text
api_rest/
├── cmd/api/                 # Ponto de entrada da aplicação Backend (main.go)
├── internal/                # Camadas e lógica interna do Go
│   ├── config/              # Variáveis de ambiente e DSN
│   ├── database/            # Conexão e migrations automáticas (GORM)
│   ├── handler/             # Handlers HTTP para processamento de JSON
│   ├── middleware/          # Interceptador de segurança (AuthRequired JWT)
│   ├── model/               # Entidades e structs (Tarefa, Usuario)
│   ├── repository/          # Camada de persistência e banco de dados
│   ├── service/             # Regras de negócio, Bcrypt e assinatura JWT
│   └── router/              # Configuração e agrupamento de rotas Gin
├── frontend/                # Interface do Usuário (TypeScript / Vite / React)
│   ├── src/                 # Componentes visuais e telas do app
│   ├── package.json         # Gerenciador de dependências do Frontend
│   └── vite.config.ts       # Configurações do servidor Vite (Porta 5173)
├── .env.example             # Modelo de configuração local do Backend
├── .gitignore               # Arquivos ignorados no versionamento
└── go.mod                   # Gerenciador de módulos e dependências do Go

```

## Pré-requisitos

- **Go** 1.26 ou superior
- **Node.js** Versão LTS recomendada
- **PostreSQL** em execução (Porta padrão `5433` ou customizável)



## Configuração

A conexão com o banco é definida por variáveis de ambiente. Em desenvolvimento, copie o arquivo de exemplo e ajuste os valores:

```bash
cp .env.example .env
```

| Variável | Descrição | Padrão |
| :--- | :--- | :--- |
| `DB_HOST` | Host do PostgreSQL | `localhost` |
| `DB_USER` | Usuário do banco | `postgres` |
| `DB_PASSWORD` | Senha do banco | `minhasenha` |
| `DB_NAME` | Nome do banco | `postgres` |
| `DB_PORT` | Porta do banco | `5433` |
| `DB_SSLMODE` | Modo SSL (`disable`, `require`, etc.) | `disable` |
| `JWT_SECRET` | Chave de assinatura dos tokens (`opcional se houver fallback`) | `variável de sistema` |

O arquivo `.env` é ignorado pelo Git. Em produção, defina as variáveis diretamente no ambiente (Docker, Kubernetes, CI, etc.) — o `.env` não é obrigatório.

Se o `.env` não existir, a API usa as variáveis do sistema operacional ou os valores padrão da tabela acima.

## Como executar

Para que a aplicação funcione por completo, o **Backend** e o **Frontend** devem ser executados simultaneamente em dois terminais separados.

1. Clone o repositório:

   ```bash
   git clone https://github.com/TrentinJr/API_Rest.git
   cd API_Rest
   ```

2. Configure o banco (veja [Configuração](#configuração)).

3. Inicializando o Back-end
  
  O servidor backend processa as regras de negócio e os endpoints, rodando por padrão na porta 8080.
  
  1. Abra um terminal na raiz do projeto (pasta API_Rest).
  2. Navegue até a pasta da API, baixe os módulos e execute o servidor:

  ```bash
    cd cmd/api
    go mod download
    go run main.go
  ```
3. Inicializando o Front-end

   A interface visual em React foi configurada para rodar de forma isolada na porta 5173 para não entrar em conflito com o Go.
  
  Abra um segundo terminal no seu VS Code (mantenha o terminal do Go aberto e rodando).

  Garanta que está na raiz do projeto e navegue até a pasta frontend:

  ```bash
  cd frontend
  ```
  Instale os pacotes necessários (apenas na primeira vez) e inicie o servidor de desenvolvimento do Vite:

  ```bash
  npm install
  npm run dev
  ```
  


A API sobe em `http://localhost:8080`. As migrations das tabelas são aplicadas automaticamente na inicialização.

## Endereço de Acesso
- Interface do Usuário (Frontend): http://localhost:5173
- Servidor de Dados (Backend API): http://localhost:8080

## 🔌 Integração & Endpoints da API (Backend)

Esta seção detalha os endpoints que a API em Go disponibiliza na porta `8080`. O Frontend consome essas rotas automaticamente através de requisições HTTP (via Fetch/Axios).

**Base URL da API:** `http://localhost:8080`

### 👤 Autenticação e Usuários (Rotas Públicas)
*Consumidas pelas telas de Login e Cadastro do Frontend.*

| Método | Endpoint | Descrição | Payload Esperado (`JSON`) |
| :--- | :--- | :--- | :--- |
| `POST` | `/usuario` | Cadastra um novo usuário | `{"email": "...", "senha": "..."}` |
| `POST` | `/usuario/login` | Autentica o usuário e gera o Token JWT | `{"email": "...", "senha": "..."}` |

### 📝 Gerenciamento de Tarefas (Rotas Protegidas por JWT)
*O Frontend gerencia o armazenamento do Token JWT no navegador e o anexa automaticamente nas requisições abaixo.*

| Método | Endpoint | Descrição | Payload Esperado (`JSON`) |
| :--- | :--- | :--- | :--- |
| `GET` | `/tarefas` | Lista as tarefas do usuário logado | *Nenhum* |
| `POST` | `/tarefas` | Cria uma nova tarefa | `{"titulo": "...", "descricao": "...", "status": "..."}` |
| `PUT` | `/tarefas/:id` | Atualiza os dados de uma tarefa | `{"titulo": "...", "descricao": "...", "status": "..."}` |
| `PATCH` | `/tarefas/:id/status`| Atualiza apenas o status (Ex: Concluído)| `{"status": "concluido"}` |
| `DELETE`| `/tarefas/:id` | Remove uma tarefa definitivamente | *Nenhum* |                   |

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
