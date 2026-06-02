package model

import "gorm.io/gorm"

type Tarefa struct {
	gorm.Model
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Concluida bool   `json:"concluida"`
}
