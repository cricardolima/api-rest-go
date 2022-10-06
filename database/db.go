package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnector() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Database connection error")
	}
}
