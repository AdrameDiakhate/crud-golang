package main

import (
	"golang-crud/controllers"
	"golang-crud/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", controllers.FindProducts)
	r.POST("/api/products", controllers.CreateProduct)
	r.GET("/api/products/one", controllers.FindProduct)
	r.PUT("/api/products/update", controllers.UpdateProduct)
	r.Run()
}
