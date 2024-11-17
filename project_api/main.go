package main

import (
	"log"
	"trucode3-challenge-final/project_api/data"
	"trucode3-challenge-final/project_api/database"
	"trucode3-challenge-final/project_api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	data.AddRoutes(router)

	return router
}

func loadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// cargar las variables de entorno
	loadEnvVars()
	// conectar a la base de datos
	db := database.CreateDbConnection()
	// se verifica si hay registros en el database
	var count int64
	db.Model(&models.DataRecord{}).Count(&count)
	if count == 0 {
		log.Println("No se encontraron datos. Iniciando conversion del dataset ...")
		database.ConvertDataSet()
	} else {
		log.Println("Los datos ya se encuentran cargados en el database")
	}

	// congigurar y ejecutar servidor
	router := setupRouter()
	router.Run("0.0.0.0:3333")

}
