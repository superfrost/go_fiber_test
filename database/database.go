package database

import (
	"fiber_test/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	connection, err := gorm.Open(sqlite.Open("images.db"), &gorm.Config{})

	if err != nil {
		panic("Can't connect to db")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
