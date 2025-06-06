package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDB() {
	dburl := os.Getenv("DB_URL")

	var err error
	DBConn, err = gorm.Open(postgres.Open(dburl), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
}
