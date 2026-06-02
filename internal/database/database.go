package database

import (
	"fmt"
	"log"

	"api_rest/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco de dados: ", err)
	}

	if err := db.AutoMigrate(&model.Tarefa{}); err != nil {
		log.Fatal("Falha ao migrar tabelas: ", err)
	}

	fmt.Println("Banco de dados conectado e tabelas criadas com sucesso!")
	return db
}
