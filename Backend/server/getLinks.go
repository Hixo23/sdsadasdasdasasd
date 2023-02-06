package server

import (
	"github.com/Hixo23/sdsadasdasdasasd/Backend/database"
	"github.com/Hixo23/sdsadasdasdasasd/Backend/models"
	"github.com/gin-gonic/gin"
)

func GetLink(c *gin.Context) {

	var link models.LinkModel
	database.DB.First(&link, c.Param("id"))
	if link.ID == 0 {
		c.JSON(404, gin.H{
			"user": "not found idiot",
		})
	}
	c.JSON(200, gin.H{
		"Link " + c.Param("id"): link,
	})

}

func GetAllLinks(c *gin.Context) {
	var links []models.LinkModel
	database.DB.Find(&links)

	c.JSON(200, gin.H{
		"All links": links,
	})

}
