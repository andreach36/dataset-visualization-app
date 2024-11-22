package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"trucode3-challenge-final/project_api/models"
)

func ExportData(filteredRecords []models.DataRecord) error {
	// verificar que haya data para exportar
	if len(filteredRecords) == 0 {
		return fmt.Errorf("no data to export")
	}
	// crear el archivo CSV
	csvFile, err := os.Create("filterData.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	// crear el escritor CSV
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// escribir la cabecera del archivo CSV
	header := []string{"Age", "Education", "Marital_status", "Occupation", "Income"}
	if err := csvWriter.Write(header); err != nil {
		return err
	}

	// escribir los registros en el archivo CSV
	for _, rowData := range filteredRecords {
		row := []string{strconv.Itoa(rowData.Age), rowData.Education, rowData.Marital_Status, rowData.Occupation, rowData.Income}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}

	return nil
}
