package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}

type Product struct {
	gorm.Model
	Title      string `json:"title" binding:"required"`
	Details    string `json:"details" binding:"required"`
	CategoryId int    `json:"category_id" binding:"required" gorm:"foreignKey:CategoryId"`
}

type UpdateProductInput struct {
	Title      string `json:"title" binding:"required"`
	Details    string `json:"details" binding:"required"`
	CategoryId int    `json:"category_id"`
}
