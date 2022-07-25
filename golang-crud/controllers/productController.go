package controllers

import (
	"net/http"
	"time"

	"golang-crud/models"

	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	title       string `json:"product_name" binding:"required"`
	description string `json:"product_detail" binding:"required"`
}

type UpdateProductInput struct {
	productName   string `json:"product_name" binding:"required"`
	productDetail string `json:"product_detail" binding:"required"`
}

func FindProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func CreateProduct(c *gin.Context) {
	// Validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create product

	product := models.Product{Title: input.title, Description: input.description, CreatedAt: time.Now()}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})

}

func FindProduct(c *gin.Context) { // Get model if exist
	var product models.Product
	id := c.Request.URL.Query().Get("id")
	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})

}

func UpdateProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	id := c.Request.URL.Query().Get("id")
	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": product})
}
