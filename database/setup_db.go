package database

import (
	"log"

	"github.com/Hixo23/sdsadasdasdasasd/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectToDb() {
	dsn := "host=localhost user=postgres password=1234 dbname=new_db port=5432 sslmode=disable TimeZone=Europe/Warsaw"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("connected to database")
	}
	DB.AutoMigrate(&models.LinkModel{})

}
