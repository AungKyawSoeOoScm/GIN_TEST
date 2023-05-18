package main

import (
	"gin_test/authentication/controllers"
	"gin_test/authentication/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	r.Run(":8080")

}
