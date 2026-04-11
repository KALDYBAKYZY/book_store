package handlers

import (
	"ginExample/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var categories = []models.Category{
	{ID: 1, Name: "Programming"},
}

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func AddCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)
	c.JSON(http.StatusCreated, newCategory)
}
