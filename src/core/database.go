package core

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DataBase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("src/files/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
