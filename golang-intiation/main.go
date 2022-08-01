package main

import (
	"golang-initiation/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/products/one", controllers.ShowOneProduct)    // Get a single product
	r.GET("/products", controllers.ShowProduct)           // Get all products
	r.POST("/products/create", controllers.CreateProduct) // Create a product
	r.PUT("/products/:id", controllers.UpdateProduct)     //Update a product
	r.DELETE("/products/:id", controllers.DeleteProduct)  //Delete a product
	r.GET("/products/search", controllers.SearchProduct)  //Search product by title

	r.Run("localhost:8080")

}
