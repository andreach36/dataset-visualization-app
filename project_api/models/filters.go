package models

import "gorm.io/gorm"

type Filter struct {
	gorm.Model     `json:"-"`
	ID             uint   `gorm:"primaryKey"`
	UserID         uint   `json:"user_id"`
	MinAge         string `json:"min_age"`
	MaxAge         string `json:"max_age"`
	Education      string `json:"education"`
	MaritalStatus  string `json:"marital_status"`
	Occupation     string `json:"occupation"`
	Income         string `json:"income"`
	OrderBy        string `json:"order_by"`
	OrderDirection string `json:"order_direction"`
}
