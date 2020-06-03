package configs

import (
	"github.com/jinzhu/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(dbType, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+DbCharset+"&parseTime=True&loc=Local")
	db.LogMode(true)
	if err != nil {
		return nil, err
	}

	return db, nil
}
