package database

import (
	"log"

	"github.com/Hixo23/sdsadasdasdasasd/Backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectToDb() {
	dsn := "host=containers-us-west-150.railway.app user=postgres password=cTjGYfdDY9piREQxKkVm dbname=railway port=7159 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("connected to database")
	}
	DB.AutoMigrate(&models.LinkModel{})

}
