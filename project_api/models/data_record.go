package models

type DataRecord struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Age            int    `csv:"age"`
	Work_Class     string `csv:"workclass"`
	Education      string `csv:"education"`
	Marital_Status string `csv:"marital-status"`
	Occupation     string `csv:"occupation"`
	Relationship   string `csv:"relationship"`
	Race           string `csv:"race"`
	Sex            string `csv:"sex"`
	Capital_Gain   int    `csv:"capital-gain"`
	Capital_Loss   int    `csv:"capital-loss"`
	Hours_Per_Week int    `csv:"hours-per-week"`
	Native_Country string `csv:"native-country"`
	Income         string `csv:"income"`
}
