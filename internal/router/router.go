package router

import (
	"api_rest/internal/handler"

	"api_rest/internal/middleware"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func Setup(tarefaHandler *handler.TarefaHandler, usuarioHandler *handler.UsuarioHandler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/usuarios/cadastrar", usuarioHandler.Cadastrar)
	r.POST("/usuarios/login", usuarioHandler.Login)

	protegido := r.Group("/")
	protegido.Use(middleware.AuthRequired()) // <-- Aplica o segurança neste grupo
	{

		protegido.POST("/tarefas", tarefaHandler.Criar)
		protegido.GET("/tarefas", tarefaHandler.Listar)
		protegido.DELETE("/tarefas/:id", tarefaHandler.Deletar)
		protegido.PATCH("/tarefas/:id/status", tarefaHandler.AtualizarStatus)
		protegido.PUT("/tarefas/:id", tarefaHandler.Editar)
	}

	return r
}
