package handler

import (
	"errors"
	"net/http"
	"strconv"

	"api_rest/internal/model"
	"api_rest/internal/service"

	"github.com/gin-gonic/gin"
)

type TarefaHandler struct {
	service *service.TarefaService
}

func NewTarefaHandler(s *service.TarefaService) *TarefaHandler {
	return &TarefaHandler{service: s}
}

func (h *TarefaHandler) Criar(c *gin.Context) {
	var novaTarefa model.Tarefa
	if err := c.ShouldBindJSON(&novaTarefa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Criar(&novaTarefa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, novaTarefa)
}

func (h *TarefaHandler) Listar(c *gin.Context) {
	tarefas, err := h.service.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tarefas)
}

func (h *TarefaHandler) Deletar(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.Deletar(id); err != nil {
		if errors.Is(err, service.ErrTarefaNaoEncontrada) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarefa deletada com sucesso!"})
}

func (h *TarefaHandler) AtualizarStatus(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var dados struct {
		Concluida bool `json:"concluida"`
	}
	if err := c.ShouldBindJSON(&dados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tarefa, err := h.service.AtualizarStatus(id, dados.Concluida)
	if err != nil {
		if errors.Is(err, service.ErrTarefaNaoEncontrada) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tarefa)
}

func (h *TarefaHandler) Editar(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var dados model.Tarefa
	if err := c.ShouldBindJSON(&dados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tarefa, err := h.service.Editar(id, &dados)
	if err != nil {
		if errors.Is(err, service.ErrTarefaNaoEncontrada) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tarefa)
}

func parseID(idParam string) (uint, error) {
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
