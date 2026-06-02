package handler

import (
	"api_rest/internal/model"
	"api_rest/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	service *service.UsuarioService
}

func NewUsuarioHandler(s *service.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{service: s}
}

// POST /usuarios/cadastrar
func (h *UsuarioHandler) Cadastrar(c *gin.Context) {
	var u model.Usuario

	// O Gin usa c.ShouldBindJSON em vez do json.NewDecoder
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if err := h.service.Cadastrar(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// POST /usuarios/login
func (h *UsuarioHandler) Login(c *gin.Context) {
	var credenciais struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	if err := c.ShouldBindJSON(&credenciais); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	token, err := h.service.Login(credenciais.Email, credenciais.Senha)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Retorna o token em formato JSON usando o Gin
	c.JSON(http.StatusOK, gin.H{"token": token})
}
