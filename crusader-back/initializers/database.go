package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func connectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failest to connect to database: ", err)
	}
}

func GetDB() *gorm.DB {

	if db == nil {
		connectToDatabase()
	}

	return db
}
