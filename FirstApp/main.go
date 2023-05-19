package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("templates/**/**")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Hello Gimn",
		})
	})
	log.Println("Server Started!")
	r.Run() // Default Port 8080
}
