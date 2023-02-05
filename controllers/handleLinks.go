package controllers

import (
	"github.com/gin-gonic/gin"
)

func CreateLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
