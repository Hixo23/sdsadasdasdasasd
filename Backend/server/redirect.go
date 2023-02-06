package server

import (
	"github.com/Hixo23/sdsadasdasdasasd/Backend/database"
	"github.com/Hixo23/sdsadasdasdasasd/Backend/models"
	"github.com/gin-gonic/gin"
)

func RedirectToLink(c *gin.Context) {
	var Link models.LinkModel
	database.DB.First(&Link, c.Param("name"))

	url := Link.Url
	c.Redirect(202, url)

}
