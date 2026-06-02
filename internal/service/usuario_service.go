package service

import (
	"api_rest/internal/model"
	"api_rest/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("chave_secreta_para_desenvolvimento") // Idealmente viria do .env

type UsuarioService struct {
	repo *repository.UsuarioRepository
}

func NewUsuarioService(repo *repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repo: repo}
}

func (s *UsuarioService) Cadastrar(u *model.Usuario) error {
	hashedSenha, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Senha = string(hashedSenha)
	return s.repo.Criar(u)
}

func (s *UsuarioService) Login(email, senha string) (string, error) {
	u, err := s.repo.BuscarPorEmail(email)
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senha))
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(JwtKey)
}

func (s *UsuarioService) Editar(u *model.Usuario) error {
	if u.Senha != "" {
		hashedSenha, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Senha = string(hashedSenha)
	}
	return s.repo.Atualizar(u)
}

func (s *UsuarioService) Apagar(id uint) error {
	return s.repo.Deletar(id)
}
