package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() {
	var err error
	Db, err = gorm.Open(sqlite.Open("src/files/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func GetDatabase() *gorm.DB {
	if Db == nil {
		InitDatabase()
	}
	return Db
}
