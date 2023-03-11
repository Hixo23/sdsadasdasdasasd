package api

import (
	"fmt"
	"net/http"

	"github.com/Hixo23/sdsadasdasdasasd/Backend/pkg/database"
	"github.com/Hixo23/sdsadasdasdasasd/Backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func HandleGetLinks(c *gin.Context) {
	var links []models.Link

	database.DB.Find(&links)
	c.JSON(200, gin.H{
		"links": links,
	})
}

func HandleCreateLink(c *gin.Context) {
	var bindLink models.BindLink
	var link models.Link

	if err := c.BindJSON(&bindLink); err != nil {
		c.String(400, "error binding link: %c", err)
	}

	database.DB.First(&link, "name = ?", bindLink.Name)

	if link.Name != "" {
		c.String(500, "this name is already in use")
	} else {
		link := models.Link{Url: bindLink.Url, Name: bindLink.Name}

		database.DB.Create(&link)
		c.JSON(202, gin.H{
			"Link created, details ": link,
		})
	}
}

func HandleUpdateLink(c *gin.Context) {
	var link models.Link

	if err := database.DB.First(&link, c.Param("id")); err != nil {
		c.String(500, "error finding Link %c", err)
	}

	if err := c.BindJSON(&link); err != nil {
		c.String(400, "binding failed %c", err)
	}

	database.DB.Model(&link).Updates(models.Link{Url: link.Url, Name: link.Name})
	c.JSON(202, gin.H{
		"link updated : ": link,
	})
}

func HandleDeleteLink(c *gin.Context) {
	var link models.Link

	database.DB.First(&link, c.Param("id"))

	database.DB.Delete(&link, c.Param("id"))

	c.JSON(204, gin.H{
		"Link": "deleted",
	})
}

func HandleRedirect(c *gin.Context) {
	var link models.Link
	name := c.Param("name")

	database.DB.First(&link, "name = ?", name)

	url := link.Url
	fmt.Println(url)
	c.Redirect(http.StatusFound, "http://"+url)
}
