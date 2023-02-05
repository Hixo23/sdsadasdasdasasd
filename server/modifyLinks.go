package server

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
	link := models.LinkModel{Url: Link1.Url, Name: Link1.Name}

	database.DB.Create(&link)

	c.JSON(202, gin.H{
		"Link created, details ": link,
	})
}
func UpdateLink(c *gin.Context) {
	var link *models.LinkModel

	database.DB.First(&link, c.Param("id"))
	c.Bind(&link)
	database.DB.Model(&link).Updates(models.LinkModel{Url: link.Url, Name: link.Name})
	c.JSON(202, gin.H{
		"link updated : ": link,
	})
}

func DeleteLink(c *gin.Context) {
	database.DB.Delete(c.Param("id"))
	c.JSON(204, gin.H{
		"Link": "deleted",
	})
}
