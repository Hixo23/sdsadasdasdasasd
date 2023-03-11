package database

import (
	"log"

	"github.com/Hixo23/sdsadasdasdasasd/Backend/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=links port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("connected to database")
	}

	DB.AutoMigrate(&models.Link{})
}
