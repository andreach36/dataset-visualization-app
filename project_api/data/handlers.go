package data

import (
	"net/http"
	"trucode3-challenge-final/project_api/database"
	"trucode3-challenge-final/project_api/models"

	"github.com/gin-gonic/gin"
)

func ShowAllData(c *gin.Context) {
	var dataRecords []models.DataRecord
	db := database.CreateDbConnection()
	tx := db.Find(&dataRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data record"})
		return
	}

	c.JSON(http.StatusOK, dataRecords)

}

func FilterDataByEducation(c *gin.Context) {
	education := c.Query("education")
	if education == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'education' query parameter"})
		return
	}
	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()
	tx := db.Where("education = ?", education).Find(&filteredRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by education"})
		return
	}

	c.JSON(http.StatusOK, filteredRecords)
}

func FilterDataByMaritalStatus(c *gin.Context) {
	maritalStatus := c.Query("marital_status")
	if maritalStatus == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'marital status' query parameter"})
		return
	}
	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()
	tx := db.Where("marital_status = ?", maritalStatus).Find(&filteredRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by marital status"})
		return
	}

	c.JSON(http.StatusOK, filteredRecords)

}

func FilterDataByOcupation(c *gin.Context) {
	occupation := c.Query("occupation")
	if occupation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'occupation' query parameter"})
		return
	}
	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()
	tx := db.Where("occupation = ?", occupation).Find(&filteredRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by occupation"})
		return
	}

	c.JSON(http.StatusOK, filteredRecords)
}

func FilterDataByAge(c *gin.Context) {
	minAge := c.Query("min_age")
	maxAge := c.Query("max_age")
	if minAge == "" || maxAge == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'education' query parameter"})
		return
	}
	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()
	tx := db.Where("age BETWEEN ? AND ?", minAge, maxAge).Find(&filteredRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by age"})
		return
	}

	c.JSON(http.StatusOK, filteredRecords)

}

func FilterDataByIncome(c *gin.Context) {
	income := c.Query("income")
	var filteredRecords []models.DataRecord
	db := database.CreateDbConnection()
	tx := db.Where("income = ?", income).Find(&filteredRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by income range"})
		return
	}

	c.JSON(http.StatusOK, filteredRecords)

}

func OrderDataByAscendant(c *gin.Context) {
	orderBy := c.Query("oder_by")
	var orderedRecords []models.DataRecord
	db := database.CreateDbConnection()
	tx := db.Order(orderBy + " ASC").Find(&orderedRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by income range"})
		return
	}

	c.JSON(http.StatusOK, orderedRecords)
}

func OrderDataByDescendant(c *gin.Context) {
	orderBy := c.Query("oder_by")
	var orderedRecords []models.DataRecord
	db := database.CreateDbConnection()
	tx := db.Order(orderBy + " DESC").Find(&orderedRecords)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter data by income range"})
		return
	}

	c.JSON(http.StatusOK, orderedRecords)

}
