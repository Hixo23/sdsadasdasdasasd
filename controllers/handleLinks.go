package controllers

import (
	"github.com/Hixo23/sdsadasdasdasasd/database"
	"github.com/Hixo23/sdsadasdasdasasd/models"
	"github.com/gin-gonic/gin"
)

func CreateLink(c *gin.Context) {
	var Link1 struct {
		Url  string `json:"Url"`
		Name string `json:"name"`
	}

	c.Bind(&Link1)
	link := models.LinkModel{Url: Link1.Url}

	database.DB.Create(&link)
	c.JSON(202, gin.H{
		"Link created, details": link,
	})
}

func GetLink(c *gin.Context) {

	var link models.LinkModel
	database.DB.First(&link, c.Param("id"))

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
