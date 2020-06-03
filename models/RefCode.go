package models

// RC struct for get value from db
type RefCode struct {
	Code        string `gorm:"column:code"`
	Message     string `gorm:"column:message"`
}

//TableName function is to get table name from db
func (RefCode) TableName() string {
	return "ref_error_codes"
}
