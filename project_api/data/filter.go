package data

import (
	"net/http"
	"trucode3-challenge-final/project_api/database"
	"trucode3-challenge-final/project_api/models"

	"github.com/gin-gonic/gin"
)

// func FilterAndOrderDataByColumns(c *gin.Context, vars ...string){
	// FilterAndOrderDataByColumns(c, "education", "occupation")
// 	for index, variable := range vars{
// 		dentre de este for voy filtrando según los campos ingresados en vars
// 	}
// }


func FilterAndOrderData(queryField string, c *gin.Context) {
	fieldValue := c.Query(queryField)
	if fieldValue == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter for" + queryField})
		return
	}
	allowedFields := map[string]bool{
		"education":      true,
		"marital_status": true,
		"occupation":     true,
	}
	if !allowedFields[queryField] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query field: " + queryField})
		return
	}

	orderBy := c.Query("order_by")
	orderDirection := c.DefaultQuery("order_direction", "ASC")
	if orderBy != "" && orderDirection != "ASC" && orderDirection != "DESC" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'order_direction' must be 'ASC' or 'DESC'"})
		return
	}

	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()

	// aplicar filtro
	query := db.Where(queryField+" = ?", fieldValue)
	// agregar ordenamiento si se especifica
	if orderBy != "" {
		query = query.Order(orderBy + " " + orderDirection)
	}
	// total de registros
	var totalRecords int64
	query.Model(&models.DataRecord{}).Count(&totalRecords)
	// aplicar pagination
	query, page, pageSize := ApplyPagination(c, query)
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