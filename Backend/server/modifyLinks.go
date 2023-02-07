package server

import (
	"net/http"

	"github.com/Hixo23/sdsadasdasdasasd/Backend/database"
	"github.com/Hixo23/sdsadasdasdasasd/Backend/models"
	"github.com/gin-gonic/gin"
)

func CreateLink(c *gin.Context) {
	var Link1 struct {
		Url  string `json:"Url" binding:"required" gorm:"unique_index"`
		Name string `json:"name" binding:"required" gorm:"unique_index"`
	}

	if err := c.ShouldBind(&Link1); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"user": "cannot create",
		})
	}
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
	var link *models.LinkModel

	database.DB.First(&link, c.Param("id"))
	database.DB.Delete(&link, c.Param("id"))

	c.JSON(204, gin.H{
		"Link": "deleted",
	})
}
