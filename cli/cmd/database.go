package cmd

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/*
function: Opens a connection to the DB
returns: database connection
 */
func OpenDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "../database.db")
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
