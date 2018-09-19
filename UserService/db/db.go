package db

import (
	"github.com/jinzhu/gorm"
)

// DB to be used
var DB *gorm.DB

// Err - error while connecting with DB
var Err error

// Connect the DB
func Connect() *gorm.DB {
	DB, Err = gorm.Open("postgres",
		"host=localhost port=5432 user=postgres dbname=user_service password=postgres sslmode=disable")
	if Err != nil {
		panic("Could not connect to DB!" + Err.Error())
	}
	return DB
}
