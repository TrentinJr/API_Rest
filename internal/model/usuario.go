package model

import "time"

type Usuario struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Senha     string    `json:"senha,omitempty" gorm:"not null"` // O omitempty ajuda a não expor a senha por engano
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
