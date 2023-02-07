package main

import (
	"github.com/Hixo23/sdsadasdasdasasd/Backend/database"
	"github.com/Hixo23/sdsadasdasdasasd/Backend/server"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	// setup database
	db := database.ConnectToDb
	db()

	// CRUD commands

	r.POST("/links", server.CreateLink)
	r.GET("/links", server.GetAllLinks)
	r.GET("/links/:id", server.GetLink)
	r.PUT("links/:id", server.UpdateLink)
	r.DELETE("links/:id", server.DeleteLink)

	//redirect
	r.GET("/links/r/:name", server.RedirectToLink)

	r.Run(":3000")

}
