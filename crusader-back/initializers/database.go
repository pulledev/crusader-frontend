package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectToDatabase() {
	var err error
	fmt.Println(os.Getenv(("DB_URL")))
	dsn := os.Getenv("DB_URL")

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
