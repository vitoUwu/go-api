package database

import (
	"log"
	"vitooapi/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func CreateConnection() {
	db, err = gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	db.AutoMigrate(&models.UserModel{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetConnection() *gorm.DB {
	return db
}
