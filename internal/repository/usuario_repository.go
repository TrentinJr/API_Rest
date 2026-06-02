package repository

import (
	"api_rest/internal/model"

	"gorm.io/gorm"
)

type UsuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{db: db}
}

func (r *UsuarioRepository) Criar(u *model.Usuario) error {
	return r.db.Create(u).Error
}

func (r *UsuarioRepository) BuscarPorEmail(email string) (*model.Usuario, error) {
	var u model.Usuario
	err := r.db.Where("email = ?", email).First(&u).Error
	return &u, err
}

func (r *UsuarioRepository) Atualizar(u *model.Usuario) error {
	return r.db.Save(u).Error
}

func (r *UsuarioRepository) Deletar(id uint) error {
	return r.db.Delete(&model.Usuario{}, id).Error
}
