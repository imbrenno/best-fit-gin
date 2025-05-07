package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imbrenno/best-fit-gin/configs"
	"github.com/imbrenno/best-fit-gin/internal/app/database"
	"github.com/imbrenno/best-fit-gin/internal/app/model"
	"github.com/imbrenno/best-fit-gin/internal/app/router"
)

func main() {
	// Carrega variáveis de ambiente
	configs.LoadEnvVariables()

	// Conecta ao banco de dados
	database.ConnectDB()
	log.Println("Banco de dados conectado com sucesso")

	// Realiza a migração automática do modelo Product
	database.DB.AutoMigrate(&model.Product{})
	log.Println("Migração do modelo Product concluída")

	// Configura o servidor e as rotas
	r := gin.Default()
	router.InitRoutes(r)
	log.Println("Rotas configuradas com sucesso")

	r.Run()
}
