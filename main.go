package main

import (
	"log"

	"github.com/ForFun21/crud-golang/src/configuration/logger"
	"github.com/ForFun21/crud-golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	logger.Info("About to start the application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}