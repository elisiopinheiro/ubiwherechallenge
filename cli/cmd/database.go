package cmd

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/*
function: Opens a connection to the DB
returns: *gorm.DB connection pointer
 */
func OpenDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "../simulator/database.db")
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
