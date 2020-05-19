package controller

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"ubiwhere/model"
)

var Db *gorm.DB

/*
function: opens the database connection to use and Migrates all the present schemas
 */
func OpenDatabase() {
	var err error
	Db, err = gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schemas
	Db.AutoMigrate(&model.SimuData{})
	Db.AutoMigrate(&model.CpuAndRam{})
}
