package main

import (
	"log"

	"api_rest/internal/config"
	"api_rest/internal/database"
	"api_rest/internal/handler"
	"api_rest/internal/model"
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

	usuarioRepo := repository.NewUsuarioRepository(db)
	usuarioService := service.NewUsuarioService(usuarioRepo)
	usuarioHandler := handler.NewUsuarioHandler(usuarioService)

	if err := db.AutoMigrate(&model.Usuario{}); err != nil {
		log.Fatalf("Erro ao executar banco de dados AutoMigrate: %v", err)
	}

	r := router.Setup(tarefaHandler, usuarioHandler)
	r.Run(":8080")
}
