package data

import (
	"net/http"
	"trucode3-challenge-final/project_api/database"
	"trucode3-challenge-final/project_api/models"

	"github.com/gin-gonic/gin"
)

func InitializeUserFilters(c *gin.Context) {
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

	// Crear una conexión a la base de datos
	db := database.CreateDbConnection()

	// Verificar si el usuario ya tiene filtros
	var filters []models.Filter
	if err := db.Where("user_id = ?", user.ID).Find(&filters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user filters"})
		return
	}

	// Si el usuario no tiene filtros, crearlos con valores predeterminados
	if len(filters) == 0 {
		defaultFilter := models.Filter{
			UserID:         user.ID,
			MinAge:         "0", // Valores predeterminados
			MaxAge:         "100",
			Education:      "",
			MaritalStatus:  "",
			Occupation:     "",
			Income:         "",
			OrderBy:        "",
			OrderDirection: "",
		}

		if err := db.Create(&defaultFilter).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create default filter"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Default filter created", "filter": defaultFilter})
		return
	}

	// Si ya tiene filtros, devolverlos
	c.JSON(http.StatusOK, gin.H{"filters": filters})
}

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

	// Obtener el ID del filtro desde la URL
	filterID := c.Param("id")

	// Verificar que el filtro pertenezca al usuario
	db := database.CreateDbConnection()
	var filter models.Filter
	if err := db.Where("id = ? AND user_id = ?", filterID, user.ID).First(&filter).Error; err != nil {
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
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Filter updated successfully", "filter": filter})
}
