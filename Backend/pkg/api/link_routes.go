package api

import (
	"github.com/Hixo23/sdsadasdasdasasd/Backend/pkg/database"
	"github.com/Hixo23/sdsadasdasdasasd/Backend/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectDB()
}

func SetupRoutes() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	r.POST("/links", HandleCreateLink)
	r.GET("/links", HandleGetLinks)
	r.PUT("links/:id", HandleUpdateLink)
	r.DELETE("links/:id", HandleDeleteLink)

	r.GET("/redirect/:name", HandleRedirect)

	r.Run(":3000")

}
