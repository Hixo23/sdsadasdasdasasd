package server

import (
	"net/http"

	"github.com/Hixo23/sdsadasdasdasasd/Backend/database"
	"github.com/Hixo23/sdsadasdasdasasd/Backend/models"
	"github.com/gin-gonic/gin"
)

func RedirectToLink(c *gin.Context) {

	var link models.LinkModel
	database.DB.Find(&link).Where(models.LinkModel{Name: c.Param("name")})
	url := link.Url

	c.Redirect(http.StatusFound, "http://"+url)
}
