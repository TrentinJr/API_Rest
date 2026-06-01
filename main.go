package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Tarefa struct {
	gorm.Model
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Concluida bool   `json:"concluida"`
}

var DB *gorm.DB

func ConectarBancoDeDados() {
	dsn := "host=localhost user=postgres password=minhasenha dbname=postgres port=5433 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco de dados: ", err)
	}

	DB.AutoMigrate(&Tarefa{})
	fmt.Println("Banco de dados conectado e tabelas criadas com sucesso!")
}

func CriarTarefa(c *gin.Context) {
	var novaTarefa Tarefa

	if err := c.ShouldBindJSON(&novaTarefa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Create(&novaTarefa)
	c.JSON(http.StatusCreated, novaTarefa)
}

func ListarTarefas(c *gin.Context) {
	var tarefas []Tarefa
	DB.Find(&tarefas)
	c.JSON(http.StatusOK, tarefas)
}

func DeletarTarefa(c *gin.Context) {
	id := c.Param("id")
	var tarefa Tarefa

	if err := DB.First(&tarefa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
		return
	}

	DB.Delete(&tarefa)
	c.JSON(http.StatusOK, gin.H{"message": "Tarefa deletada com sucesso!"})
}

func AtualizarStatusTarefa(c *gin.Context) {
	id := c.Param("id")
	var tarefa Tarefa

	if err := DB.First(&tarefa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
		return
	}

	var dadosEnviados struct {
		Concluida bool `json:"concluida"`
	}

	if err := c.ShouldBindJSON(&dadosEnviados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&tarefa).Update("Concluida", dadosEnviados.Concluida)
	c.JSON(http.StatusOK, tarefa)
}

func EditarTarefa(c *gin.Context) {
	id := c.Param("id")
	var tarefa Tarefa

	if err := DB.First(&tarefa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&tarefa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Save(&tarefa)
	c.JSON(http.StatusOK, tarefa)
}

func main() {
	ConectarBancoDeDados()

	r := gin.Default()

	r.POST("/tarefas", CriarTarefa)
	r.GET("/tarefas", ListarTarefas)
	r.DELETE("/tarefas/:id", DeletarTarefa)
	r.PATCH("/tarefas/:id/status", AtualizarStatusTarefa)
	r.PUT("/tarefas/:id", EditarTarefa)

	r.Run(":8080")
}
