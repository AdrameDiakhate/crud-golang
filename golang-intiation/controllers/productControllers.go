package controllers

import (
	"golang-initiation/config"
	"golang-initiation/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get all products in the database

func ShowProduct(c *gin.Context) {
	var products []models.Product

	db := config.ConnectToDatabase()

	db.Select("*").Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

//Creating a new product with his category
func CreateProduct(c *gin.Context) {
	db := config.ConnectToDatabase()

	var product models.Product

	error := c.ShouldBindJSON(&product)
	if error != nil {
		c.Errors.Errors()
	}
	product = models.Product{
		Title:      product.Title,
		Details:    product.Details,
		CategoryId: product.CategoryId,
	}
	db.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})

}

// Get a single product by his ID

func ShowOneProduct(c *gin.Context) {
	db := config.ConnectToDatabase()
	var product models.Product
	id := c.Request.URL.Query().Get("id")
	if err := db.Find(&product).Where("id=?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Produit indisponible!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// Update a product

func UpdateProduct(c *gin.Context) {
	//Get the product to updatec
	db := config.ConnectToDatabase()
	var product models.Product
	id := c.Request.URL.Query().Get("id")
	err := db.Find(&product).Where("id=?", id)
	if err != nil {
		c.Errors.Errors()
	}
	c.ShouldBindJSON(&product)
	db.Save(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

//Delete a product
func DeleteProduct(c *gin.Context) {
	db := config.ConnectToDatabase()
	var product models.Product
	id := c.Params.ByName("id")

	db.Where("id=?", id).Delete(&product)
	c.JSON(200, gin.H{"data": id})
}

//Search a product by title
func SearchProduct(c *gin.Context) {
	db := config.ConnectToDatabase()
	var product models.Product
	title := c.Request.URL.Query().Get("title")

	if err := db.Where("title=?", title).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Produit indisponible!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})

}
