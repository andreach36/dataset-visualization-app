package data

import (
	"net/http"
	"net/url"
	"strings"
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

// filtrar y ordenar según un campo especifico
func FilterDataByEducation(c *gin.Context) {
	FilterAndOrderData("education", c)
}

func FilterDataByMaritalStatus(c *gin.Context) {
	FilterAndOrderData("marital_status", c)
}

func FilterDataByOccupation(c *gin.Context) {
	FilterAndOrderData("occupation", c)
}

// filtrar y ordenar según un rango de edad
func FilterDataByAgeRange(c *gin.Context) {
	minAge := c.Query("min_age")
	maxAge := c.Query("max_age")
	if minAge == "" || maxAge == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'age' query parameter"})
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
	query := db.Where("age BETWEEN ? AND ?", minAge, maxAge)

	// Agregar ordenamiento si se especifica
	if orderBy != "" {
		query = query.Order(orderBy + " " + orderDirection)
	}

	// total de registros
	var totalRecords int64
	query.Model(&models.DataRecord{}).Count(&totalRecords)

	query, page, pageSize := ApplyPagination(c, query)
	tx := query.Find(&filteredRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by age"})
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

// filtrar y ordenar según los ingresos
func FilterDataByIncome(c *gin.Context) {
	income := c.Query("income")
	income, err := url.QueryUnescape(income)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if income == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'income' query parameter"})
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

	query := db.Model(&models.DataRecord{})

	// aplicar filtros según rango de ingresos
	switch {
	case strings.HasPrefix(income, "<"):
		query = db.Where("income = ?", income)
	case strings.HasPrefix(income, ">"):
		query = db.Where("income = ?", income)
	}

	// Agregar ordenamiento si se especifica
	if orderBy != "" {
		query = query.Order(orderBy + " " + orderDirection)
	}

	var totalRecords int64
	query.Model(&models.DataRecord{}).Count(&totalRecords)

	query, page, pageSize := ApplyPagination(c, query)
	tx := query.Find(&filteredRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by income range"})
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

// ordenar data en general según campo específico [sin filtrar]
func OrderData(c *gin.Context) {
	orderBy := c.Query("order_by")
	orderDirection := c.DefaultQuery("order_direction", "ASC")
	if orderBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'order_by' query parameter"})
		return
	}
	if orderDirection != "ASC" && orderDirection != "DESC" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'order_direction' must be 'ASC' or 'DESC'"})
		return
	}
	var orderedRecords []models.DataRecord
	db := database.CreateDbConnection()

	// aplicar orden
	query := db.Order(orderBy + " " + orderDirection)

	// total de registros
	var totalRecords int64
	query.Model(&models.DataRecord{}).Count(&totalRecords)

	// aplicar paginación
	query, page, pageSize := ApplyPagination(c, query)
	// obtener registros ordenados
	tx := query.Find(&orderedRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to order data by ascendent"})
		return
	}

	// responder con datos y metadatos de paginación
	c.JSON(http.StatusOK, gin.H{
		"data": orderedRecords,
		"meta": gin.H{
			"total_records": totalRecords,
			"page":          page,
			"page_size":     pageSize,
			"total_pages":   (totalRecords + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}
