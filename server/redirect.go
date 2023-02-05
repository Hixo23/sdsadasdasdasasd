package server

import (
	"github.com/Hixo23/sdsadasdasdasasd/database"
	"github.com/Hixo23/sdsadasdasdasasd/models"
	"github.com/gin-gonic/gin"
)

func RedirectLink(c *gin.Context) {
	var link models.LinkModel
	database.DB.First(&link, c.Param("id"))

	linkurl := link.Url
	c.Redirect(200, linkurl)
}
