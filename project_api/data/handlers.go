package data

import (
	"net/http"
	"strconv"
	"strings"
	"trucode3-challenge-final/project_api/database"
	"trucode3-challenge-final/project_api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApplyPagination(c *gin.Context, query *gorm.DB) (*gorm.DB, int, int) {
	// Leer parámetros de paginación
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	// convertir parámetros a enteros
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeNum < 1 {
		pageSizeNum = 10
	}

	// calcular offfset
	offset := (pageNum - 1) * pageSizeNum

	// aplicar límites de paginación
	query = query.Offset(offset).Limit(pageSizeNum)

	return query, pageNum, pageSizeNum
}

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

func FilterDataByEducation(c *gin.Context) {
	education := c.Query("education")
	if education == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'education' query parameter"})
		return
	}
	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()

	// aplicar filtro y paginación
	query := db.Where("education = ?", education)
	// total de registros
	var totalRecords int64
	query.Model(&models.DataRecord{}).Count(&totalRecords)

	query, page, pageSize := ApplyPagination(c, query)
	tx := query.Find(&filteredRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by education"})
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

func FilterDataByMaritalStatus(c *gin.Context) {
	maritalStatus := c.Query("marital_status")
	if maritalStatus == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'marital status' query parameter"})
		return
	}
	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()

	// aplicar filtro y paginación
	query := db.Where("marital_status = ?", maritalStatus)
	// total de registros
	var totalRecords int64
	query.Model(&models.DataRecord{}).Count(&totalRecords)

	query, page, pageSize := ApplyPagination(c, query)
	tx := query.Find(&filteredRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by marital status"})
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

func FilterDataByOcupation(c *gin.Context) {
	occupation := c.Query("occupation")
	if occupation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'occupation' query parameter"})
		return
	}
	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()

	// aplicar filtro y paginación
	query := db.Where("occupation = ?", occupation)
	// total de registros
	var totalRecords int64
	query.Model(&models.DataRecord{}).Count(&totalRecords)

	query, page, pageSize := ApplyPagination(c, query)
	tx := query.Find(&filteredRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by occupation"})
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

func FilterDataByAge(c *gin.Context) {
	minAge := c.Query("min_age")
	maxAge := c.Query("max_age")
	if minAge == "" || maxAge == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'age' query parameter"})
		return
	}
	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()

	// aplicar filtro y paginación
	query := db.Where("age BETWEEN ? AND ?", minAge, maxAge)
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

func FilterDataByIncome(c *gin.Context) {
	income := c.Query("income")
	if income == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'income' query parameter"})
		return
	}
	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()

	// aplicar el filtro
	query := filterIncome(income, db)
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

func filterIncome(income string, db *gorm.DB) *gorm.DB {
	var query *gorm.DB
	if strings.HasPrefix(income, "<=") {
		query = db.Where("income <= ?", strings.TrimPrefix(income, "<="))
	} else if strings.HasPrefix(income, ">=") {
		query = db.Where("income >= ?", strings.TrimPrefix(income, ">="))
	} else if strings.HasPrefix(income, "<") {
		query = db.Where("income < ?", strings.TrimPrefix(income, "<"))
	} else if strings.HasPrefix(income, ">") {
		query = db.Where("income > ?", strings.TrimPrefix(income, ">"))
	}
	return query
}

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
