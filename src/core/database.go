package core

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("src/files/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func GetDatabase() *gorm.DB {
	if db == nil {
		InitDatabase()
	}
	return db
}
