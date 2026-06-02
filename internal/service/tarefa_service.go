package service

import (
	"errors"

	"api_rest/internal/model"
	"api_rest/internal/repository"

	"gorm.io/gorm"
)

var ErrTarefaNaoEncontrada = errors.New("tarefa não encontrada")

type TarefaService struct {
	repo repository.TarefaRepository
}

func NewTarefaService(repo repository.TarefaRepository) *TarefaService {
	return &TarefaService{repo: repo}
}

func (s *TarefaService) Criar(tarefa *model.Tarefa) error {
	return s.repo.Create(tarefa)
}

func (s *TarefaService) Listar() ([]model.Tarefa, error) {
	return s.repo.FindAll()
}

func (s *TarefaService) BuscarPorID(id uint) (*model.Tarefa, error) {
	tarefa, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTarefaNaoEncontrada
		}
		return nil, err
	}
	return tarefa, nil
}

func (s *TarefaService) Deletar(id uint) error {
	tarefa, err := s.BuscarPorID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(tarefa)
}

func (s *TarefaService) AtualizarStatus(id uint, concluida bool) (*model.Tarefa, error) {
	if _, err := s.BuscarPorID(id); err != nil {
		return nil, err
	}
	if err := s.repo.UpdateStatus(id, concluida); err != nil {
		return nil, err
	}
	tarefa, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return tarefa, nil
}

func (s *TarefaService) Editar(id uint, dados *model.Tarefa) (*model.Tarefa, error) {
	tarefa, err := s.BuscarPorID(id)
	if err != nil {
		return nil, err
	}

	tarefa.Titulo = dados.Titulo
	tarefa.Descricao = dados.Descricao

	if err := s.repo.Save(tarefa); err != nil {
		return nil, err
	}
	return tarefa, nil
}
