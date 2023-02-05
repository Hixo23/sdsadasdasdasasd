package main

import (
	"github.com/Hixo23/sdsadasdasdasasd/database"
	"github.com/Hixo23/sdsadasdasdasasd/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// setup database
	db := database.ConnectToDb
	db()
	// CRUD commands
	r.POST("/links", server.CreateLink)
	r.GET("/links", server.GetAllLinks)
	r.GET("/links/:id", server.GetLink)
	r.GET("/links/redirect/:id", server.RedirectLink)
	r.PUT("links/:id", server.UpdateLink)
	r.DELETE("links/:id", server.DeleteLink)
	r.Run(":3000")

}
