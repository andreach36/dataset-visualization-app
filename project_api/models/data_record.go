package models

type DataRecord struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Age            int    `csv:"age"`
	Education      string `csv:"education"`
	Marital_Status string `csv:"marital-status"`
	Occupation     string `csv:"occupation"`
	Income         string `csv:"income"`
}
