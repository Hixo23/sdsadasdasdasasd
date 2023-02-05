package main

import (
	"github.com/Hixo23/sdsadasdasdasasd/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.CreateLink)
	r.Run(":3000")
}
