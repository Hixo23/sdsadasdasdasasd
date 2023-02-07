package server

import (
	"github.com/Hixo23/sdsadasdasdasasd/Backend/database"
	"github.com/Hixo23/sdsadasdasdasasd/Backend/models"
	"github.com/gin-gonic/gin"
)

func GetLink(c *gin.Context) {

	var link *models.LinkModel

	database.DB.First(&link, c.Param("id"))

	if link.ID == 0 && link.Name == "" {
		c.JSON(404, gin.H{
			"link": "not found",
		})
	} else {

		c.JSON(200, gin.H{
			"Link " + c.Param("id"): link,
		})
	}

}

func GetAllLinks(c *gin.Context) {

	var links *[]models.LinkModel

	database.DB.Find(&links)

	c.JSON(200, gin.H{
		"links": links,
	})

}
