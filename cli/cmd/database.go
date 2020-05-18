package cmd

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func OpenDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "../simulator/database.db")
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
