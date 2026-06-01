# 🚀 Go Task Manager API

Uma API REST funcional e robusta para gerenciamento de tarefas (To-Do List), desenvolvida em **Go (Golang)**. O projeto utiliza práticas modernas de desenvolvimento back-end, incluindo mapeamento objeto-relacional (ORM) e validação de rotas.

## 🛠️ Tecnologias Utilizadas

* **Linguagem:** Go (Golang)
* **Framework Web:** [Gin Gonic](https://github.com/gin-gonic/gin) (Roteamento de alta performance)
* **ORM:** [GORM](https://gorm.io/) (Interação simplificada com o banco de dados)
* **Banco de Dados:** PostgreSQL
* **Testes de Endpoints:** Postman

## 📌 Rotas da API (Endpoints)

A API gerencia os recursos de tarefas através dos seguintes endpoints expostos na porta `:8080`:

| Método | Endpoint | Descrição | Corpo da Requisição (JSON) |
| :--- | :--- | :--- | :--- |
| **POST** | `/tarefas` | Cria uma nova tarefa | `{"titulo": "Estudar Go", "descricao": "Praticar GORM"}` |
| **GET** | `/tarefas` | Lista todas as tarefas cadastradas | Nenhum |
| **PUT** | `/tarefas/:id` | Edita o título e a descrição de uma tarefa | `{"titulo": "Novo Título", "descricao": "Nova Descrição"}` |
| **PATCH** | `/tarefas/:id/status` | Altera apenas o status de conclusão | `{"concluida": true}` |
| **DELETE** | `/tarefas/:id` | Deleta uma tarefa pelo ID | Nenhum |

## 🏁 Como Executar o Projeto Localmente

### Pré-requisitos
* Ter o **Go** instalado em sua máquina.
* Ter um banco de dados **PostgreSQL** rodando (na porta `5433` ou ajustar a string de conexão no código).

### Passo a Passo

1. Clone o repositório:
   ```bash
   git clone [https://github.com/seu-usuario/nome-do-seu-repositorio.git](https://github.com/seu-usuario/nome-do-seu-repositorio.git)