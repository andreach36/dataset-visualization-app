package main

import (
	"log"
	"trucode3-challenge-final/project_api/database"

	"github.com/joho/godotenv"
)

func loadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// cargar las variables de entorno
	loadEnvVars()
	// convertir el DataSet
	database.ConvertDataSet()

}
