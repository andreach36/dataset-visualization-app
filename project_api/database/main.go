package database

import (
	"fmt"
	"log"
	"os"
	"trucode3-challenge-final/project_api/models"

	"github.com/gocarina/gocsv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDbConnection() *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ConvertDataSet() {

	// conectar con la base de datos
	db := CreateDbConnection()

	// Crear el tipo enum en PostgreSQL
	db.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum') THEN CREATE TYPE status_enum AS ENUM ('active', 'inactive', 'pending'); END IF; END $$;")

	// Migrar el esquema para crear la tabla
	if err := db.AutoMigrate(&models.UserRecord{}); err != nil {
		log.Fatalf("Error en la migración de la base de datos: " + err.Error())
	}

	// open the csv file
	file, err := os.Open("../data/source.data")

	if err != nil {
		log.Fatalf("Error al abrir el archivo csv" + err.Error())
	}

	defer file.Close()

	// read the csv file into a slice of record structs
	var userRecords []models.UserRecord
	if err := gocsv.UnmarshalFile(file, &userRecords); err != nil {
		log.Fatalf("Error al convertir csv a datos estructurados" + err.Error())
	}

	// procesar los registros en lotes concurrentes
	batchUserRecord(db, userRecords)

}

func saveUserRecord(db *gorm.DB, userRecords []models.UserRecord, done chan<- bool) {
	// guardar los registros en la base de datos
	if err := db.Create(&userRecords).Error; err != nil {
		log.Printf("Error al guardar el lote de registros: %v", err)
		done <- false
		return
	}

	done <- true
	fmt.Println("Lote de datos importado correctamente en la base de datos")
}

func batchUserRecord(db *gorm.DB, userRecords []models.UserRecord) {
	const batchSize = 1000 // numero de registros por lote
	done := make(chan bool)
	totalBatches := (len(userRecords) + batchSize - 1) / batchSize

	// procesar cada lote en una goroutine
	for i := 0; i < len(userRecords); i += batchSize {
		end := i + batchSize
		if end > len(userRecords) {
			end = len(userRecords)
		}
		// lanzar una goroutine para cada lote
		go saveUserRecord(db, userRecords[i:end], done)
	}
	// esperar a que todos los lotes terminen
	for i := 0; i < totalBatches; i++ {
		if succes := <-done; !succes {
			fmt.Println("Error al insertar algunos registros")
		}
	}

	close(done) // cerrar el canal al final

}
