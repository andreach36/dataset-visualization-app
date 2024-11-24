package data

import (
	"bytes"
	"encoding/csv"
	"net/http"
	"strconv"
	"trucode3-challenge-final/project_api/models"

	"github.com/gin-gonic/gin"
)

func ExportData(c *gin.Context, filteredRecords []models.DataRecord) {
	// verificar que haya data para exportar
	if len(filteredRecords) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No data to export"})
		return
	}
	// Crear el archivo CSV en memoria
	var buf bytes.Buffer
	csvWriter := csv.NewWriter(&buf)

	// escribir la cabecera del archivo CSV
	header := []string{"Age", "Work_Class", "Education", "Marital_status", "Occupation", "Relationship", "Race", "Sex", "Capital_Gain", "Capital_Loss", "Hours_per_week", "Native_Country", "Income"}
	if err := csvWriter.Write(header); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write header to CSV"})
		return
	}

	// escribir los registros en el archivo CSV
	for _, rowData := range filteredRecords {
		row := []string{strconv.Itoa(rowData.Age), rowData.Work_Class, rowData.Education, rowData.Marital_Status, rowData.Occupation, rowData.Relationship, rowData.Race, rowData.Sex, strconv.Itoa(rowData.Capital_Gain), strconv.Itoa(rowData.Capital_Loss), strconv.Itoa(rowData.Hours_Per_Week), rowData.Native_Country, rowData.Income}
		if err := csvWriter.Write(row); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write row"})
			return
		}
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write CSV data"})
		return
	}

	// Configurar los encabezados para la descarga del archivo
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=filteredData.csv")
	c.Status(http.StatusOK)

	// Enviar el archivo CSV como respuesta
	c.Data(http.StatusOK, "application/octet-stream", buf.Bytes())

}
