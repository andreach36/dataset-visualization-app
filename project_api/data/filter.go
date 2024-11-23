package data

import (
	"fmt"
	"net/http"
	"trucode3-challenge-final/project_api/database"
	"trucode3-challenge-final/project_api/models"

	"github.com/gin-gonic/gin"
)

func UpdateUserFilter(c *gin.Context) {
	// Obtener el usuario autenticado del contexto
	authUser, exists := c.Get("authorizedUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, ok := authUser.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve authenticated user"})
		return
	}

	// Verificar que el filtro pertenezca al usuario
	db := database.CreateDbConnection()
	var filter models.Filter
	if err := db.Where("user_id = ?", user.ID).First(&filter).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Filter not found or not owned by user"})
		return
	}

	// Actualizar los campos del filtro con los nuevos valores
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.Save(&filter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update filter"})
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Filter updated successfully", "filter": filter})
}

func GetUserFilter(c *gin.Context) {
	// Obtener el usuario autenticado del contexto
	authUser, exists := c.Get("authorizedUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, ok := authUser.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve authenticated user"})
		return
	}
	// Verificar que el filtro pertenezca al usuario
	db := database.CreateDbConnection()
	var filter models.Filter
	if err := db.Where("user_id = ?", user.ID).First(&filter).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Filter not found or not owned by user"})
		return
	}
	tx := db.Find(&filter)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get filter configuration"})
		return
	}
	// responder con datos
	c.JSON(http.StatusOK, gin.H{
		"filters": filter,
	})

}
