package main

import (
	"log"

	"api_rest/internal/config"
	"api_rest/internal/database"
	"api_rest/internal/handler"
	"api_rest/internal/repository"
	"api_rest/internal/router"
	"api_rest/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	db := database.Connect(config.LoadDatabase().DSN())

	tarefaRepo := repository.NewTarefaRepository(db)
	tarefaService := service.NewTarefaService(tarefaRepo)
	tarefaHandler := handler.NewTarefaHandler(tarefaService)

	r := router.Setup(tarefaHandler)
	r.Run(":8080")
}
