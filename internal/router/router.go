package router

import (
	"api_rest/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup(tarefaHandler *handler.TarefaHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/tarefas", tarefaHandler.Criar)
	r.GET("/tarefas", tarefaHandler.Listar)
	r.DELETE("/tarefas/:id", tarefaHandler.Deletar)
	r.PATCH("/tarefas/:id/status", tarefaHandler.AtualizarStatus)
	r.PUT("/tarefas/:id", tarefaHandler.Editar)

	return r
}
