package database

import (
	"fmt"
	"log"
	"os"
	"strings"
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

	// Migrar el esquema para crear la tabla
	if err := db.AutoMigrate(&models.DataRecord{}); err != nil {
		log.Fatalf("Error en la migración de la base de datos: " + err.Error())
	}

	// open the csv file
	file, err := os.Open("../data/source.data")

	if err != nil {
		log.Fatalf("Error al abrir el archivo csv" + err.Error())
	}

	defer file.Close()

	// read the csv file into a slice of record structs
	var dataRecords []models.DataRecord
	if err := gocsv.UnmarshalFile(file, &dataRecords); err != nil {
		log.Fatalf("Error al convertir csv a datos estructurados" + err.Error())
	}

	// limpiar los espacios en blanco de los datos
	for i := range dataRecords {
		dataRecords[i].Education = strings.TrimSpace(dataRecords[i].Education)
		dataRecords[i].Marital_Status = strings.TrimSpace(dataRecords[i].Marital_Status)
		dataRecords[i].Occupation = strings.TrimSpace(dataRecords[i].Occupation)
		dataRecords[i].Income = strings.TrimSpace(dataRecords[i].Income)
	}

	// procesar los registros en lotes concurrentes
	batchUserRecord(db, dataRecords)

}

func saveUserRecord(db *gorm.DB, dataRecords []models.DataRecord, done chan<- bool) {
	// guardar los registros en la base de datos
	if err := db.Create(&dataRecords).Error; err != nil {
		log.Printf("Error al guardar el lote de registros: %v", err)
		done <- false
		return
	}

	done <- true
	fmt.Println("Lote de datos importado correctamente en la base de datos")
}

func batchUserRecord(db *gorm.DB, dataRecords []models.DataRecord) {
	const batchSize = 1000 // numero de registros por lote
	done := make(chan bool)
	totalBatches := (len(dataRecords) + batchSize - 1) / batchSize

	// procesar cada lote en una goroutine
	for i := 0; i < len(dataRecords); i += batchSize {
		end := i + batchSize
		if end > len(dataRecords) {
			end = len(dataRecords)
		}
		// lanzar una goroutine para cada lote
		go saveUserRecord(db, dataRecords[i:end], done)
	}
	// esperar a que todos los lotes terminen
	for i := 0; i < totalBatches; i++ {
		if succes := <-done; !succes {
			fmt.Println("Error al insertar algunos registros")
		}
	}

	close(done) // cerrar el canal al final

}
