package data

import (
	"strconv"

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
