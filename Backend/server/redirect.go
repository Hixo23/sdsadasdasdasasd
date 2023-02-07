package server

import (
	"net/http"

	"github.com/Hixo23/sdsadasdasdasasd/Backend/database"
	"github.com/Hixo23/sdsadasdasdasasd/Backend/models"
	"github.com/gin-gonic/gin"
)

func RedirectToLink(c *gin.Context) {
	var link models.LinkModel
	database.DB.First(&link, "name = ?", c.Param("name"))

	url := link.Url
	if url == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"link": "not found",
		})
	} else {

		c.Redirect(http.StatusFound, "http://"+url)

	}
}
