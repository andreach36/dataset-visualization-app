package main

import (
	"log"
	"os"
	"project_api/auth"
	"project_api/data"
	"project_api/database"
	"project_api/models"
	"project_api/shared"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(shared.Cors())
	auth.AddRoutes(router)
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
	if os.Getenv("GIN_MODE") != "release" {
		// cargar las variables de entorno
		loadEnvVars()
	}
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

	// inicializar la tabla de usuario y filtros de usuario
	db.AutoMigrate(&models.User{}, &models.Filter{})

	// congigurar y ejecutar servidor
	router := setupRouter()
	router.Run(os.Getenv("PORT"))

}
