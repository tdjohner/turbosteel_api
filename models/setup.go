// models/setup.go

package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("dogdogdog.db"), &gorm.Config{})

	if err != nil {
		panic("connect to database failed")
	}

	err = database.AutoMigrate(&Dog{})
	if err != nil {
		return
	}

	DB = database
}
