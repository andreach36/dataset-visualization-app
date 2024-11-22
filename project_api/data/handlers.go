package data

import (
	"net/http"
	"net/url"

	"trucode3-challenge-final/project_api/database"
	"trucode3-challenge-final/project_api/models"

	"github.com/gin-gonic/gin"
)

// mostrar todos los datos
func ShowAllData(c *gin.Context) {
	var dataRecords []models.DataRecord
	db := database.CreateDbConnection()

	// aplicar paginación
	query := db.Model(&models.DataRecord{})
	query, page, pageSize := ApplyPagination(c, query)

	// obtener datos
	tx := query.Find(&dataRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data record"})
		return
	}

	// total de registros
	var totalRecords int64
	db.Model(&models.DataRecord{}).Count(&totalRecords)

	// responder con datos y metadatos de paginación
	c.JSON(http.StatusOK, gin.H{
		"data": dataRecords,
		"meta": gin.H{
			"total_records": totalRecords,
			"page":          page,
			"page_size":     pageSize,
			"total_pages":   (totalRecords + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

func FilterAndOrderData(c *gin.Context) {
	// Recoger los parámetros de la consulta
	minAge := c.DefaultQuery("min_age", "0")
	maxAge := c.DefaultQuery("max_age", "10")
	education := c.Query("education")
	maritalStatus := c.Query("marital_status")
	occupation := c.Query("occupation")
	income, _ := url.QueryUnescape(c.Query("income"))

	// Validación de campos permitidos
	allowedFields := map[string]bool{
		"education":      true,
		"marital_status": true,
		"occupation":     true,
		"income":         true,
	}

	// Validación para el rango de edad
	if minAge == "" || maxAge == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'age' query parameter"})
		return
	}

	// Validación filtrado de campos
	if education != "" && !allowedFields["education"] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter field: education"})
		return
	}
	if maritalStatus != "" && !allowedFields["marital_status"] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter field: marital_status"})
		return
	}
	if occupation != "" && !allowedFields["occupation"] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter field: occupation"})
		return
	}
	if income != "" && !allowedFields["income"] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter field: income"})
		return
	}

	// construir la consulta
	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()
	query := db.Model(&models.DataRecord{})

	// Construir filtros
	// Rango de edad
	if minAge != "" && maxAge != "" {
		query = db.Where("age BETWEEN ? AND ?", minAge, maxAge)
	}
	// educacion
	if education != "" {
		query = query.Where("education = ?", education)
	}
	// estado civil
	if maritalStatus != "" {
		query = query.Where("marital_status = ?", maritalStatus)
	}
	// ocupacion
	if occupation != "" {
		query = query.Where("occupation = ?", occupation)
	}
	// ingresos
	if income != "" {
		query = query.Where("income = ?", income) // Modifica según el formato de ingreso
	}

	// Obtener el campo de ordenamiento y la dirección
	orderBy := c.Query("order_by")
	orderDirection := c.DefaultQuery("order_direction", "ASC")
	if orderBy != "" {
		query = query.Order(orderBy + " " + orderDirection)
	}

	// total de registros
	var totalRecords int64
	query.Model(&models.DataRecord{}).Count(&totalRecords)

	// verificar si se solicita exportar los datos
	export := c.DefaultQuery("export", "false")
	if export == "true" {
		// ejecutar la consulta final
		tx := query.Find(&filteredRecords)
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data"})
			return
		}
		// exportar la data
		err := ExportData(filteredRecords)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to export data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Data exported successfully"})
		return
	} else {

		// aplicar paginacion
		query, page, pageSize := ApplyPagination(c, query)

		// ejecutar la consulta final
		tx := query.Find(&filteredRecords)
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data"})
			return
		}

		// responder con datos y metadatos de paginación
		c.JSON(http.StatusOK, gin.H{
			"data": filteredRecords,
			"meta": gin.H{
				"total_records": totalRecords,
				"page":          page,
				"page_size":     pageSize,
				"total_pages":   (totalRecords + int64(pageSize) - 1) / int64(pageSize),
			},
		})

	}

}
