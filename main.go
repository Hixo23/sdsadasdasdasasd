package main

import (
	"github.com/Hixo23/sdsadasdasdasasd/controllers"
	"github.com/Hixo23/sdsadasdasdasasd/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// setup database
	db := database.ConnectToDb
	db()
	// CRUD commands
	r.POST("/links", controllers.CreateLink)
	r.GET("/links", controllers.GetAllLinks)
	r.GET("/links/:id", controllers.GetLink)

	r.Run(":3000")

}
