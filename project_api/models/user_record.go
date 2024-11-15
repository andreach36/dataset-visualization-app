package models

type UserRecord struct {
	Status         Status `gorm:"type:status_enum;default:'inactive'"`
	Age            int    `csv:"age"`
	Workclass      string `csv:"workclass"`
	FinalWeight    int    `csv:"fnlwgt"`
	Education      string `csv:"education"`
	Education_Num  int    `csv:"education-num"`
	Marital_Status string `csv:"marital-status"`
	Occupation     string `csv:"occupation"`
	Relationship   string `csv:"relationship"`
	Race           string `csv:"race"`
	Sex            string `csv:"sex"`
	Capital_Gain   int    `csv:"capital-gain"`
	Capital_Loss   int    `csv:"capital-loss"`
	Hours_Per_Week int    `csv:"hours-per-week"`
	Native_Country string `csv:"native-country"`
	Income_Range   string `csv:"income-range"`
}

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
	Pending  Status = "pending"
)
