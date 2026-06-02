package main

import (
	"api_rest/internal/database"
	"api_rest/internal/handler"
	"api_rest/internal/repository"
	"api_rest/internal/router"
	"api_rest/internal/service"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db := database.Connect(dsn)

	tarefaRepo := repository.NewTarefaRepository(db)
	tarefaService := service.NewTarefaService(tarefaRepo)
	tarefaHandler := handler.NewTarefaHandler(tarefaService)

	r := router.Setup(tarefaHandler)
	r.Run(":8080")
}
