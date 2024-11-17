package models

type DataRecord struct {
	Status         Status `gorm:"type:status_enum;default:'inactive'"`
	Age            int    `csv:"age"`
	Education      string `csv:"education"`
	Marital_Status string `csv:"marital-status"`
	Occupation     string `csv:"occupation"`
	Income         string `csv:"income"`
}

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
	Pending  Status = "pending"
)
