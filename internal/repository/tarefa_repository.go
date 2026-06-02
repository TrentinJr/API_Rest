package repository

import (
	"api_rest/internal/model"

	"gorm.io/gorm"
)

type TarefaRepository interface {
	Create(tarefa *model.Tarefa) error
	FindAll() ([]model.Tarefa, error)
	FindByID(id uint) (*model.Tarefa, error)
	Delete(tarefa *model.Tarefa) error
	Save(tarefa *model.Tarefa) error
	UpdateStatus(id uint, concluida bool) error
}

type tarefaRepository struct {
	db *gorm.DB
}

func NewTarefaRepository(db *gorm.DB) TarefaRepository {
	return &tarefaRepository{db: db}
}

func (r *tarefaRepository) Create(tarefa *model.Tarefa) error {
	return r.db.Create(tarefa).Error
}

func (r *tarefaRepository) FindAll() ([]model.Tarefa, error) {
	var tarefas []model.Tarefa
	err := r.db.Find(&tarefas).Error
	return tarefas, err
}

func (r *tarefaRepository) FindByID(id uint) (*model.Tarefa, error) {
	var tarefa model.Tarefa
	err := r.db.First(&tarefa, id).Error
	if err != nil {
		return nil, err
	}
	return &tarefa, nil
}

func (r *tarefaRepository) Delete(tarefa *model.Tarefa) error {
	return r.db.Delete(tarefa).Error
}

func (r *tarefaRepository) Save(tarefa *model.Tarefa) error {
	return r.db.Save(tarefa).Error
}

func (r *tarefaRepository) UpdateStatus(id uint, concluida bool) error {
	return r.db.Model(&model.Tarefa{}).Where("id = ?", id).Update("concluida", concluida).Error
}
