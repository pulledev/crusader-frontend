package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectToDatabase() {
	var err error

	LoadEnvVars()
	dsn := os.Getenv("DB_URL")

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}

func GetDB() *gorm.DB {

	if db == nil {
		connectToDatabase()
	}

	return db
}
