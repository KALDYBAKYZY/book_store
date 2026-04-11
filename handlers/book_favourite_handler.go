package handlers

import (
	"ginExample/config"
	"ginExample/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
	userID := c.GetUint("user_id")
	var books []models.Book
	var total int64
	config.DB.Table("favorite_books").
		Where("user_id = ?", userID).
		Count(&total)

	err := config.DB.Table("books").
		Joins("JOIN favorite_books ON favorite_books.book_id = books.id").
		Where("favorite_books.user_id = ?", userID).
		Find(&books).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  books,
		"total": total,
	})
}

func AddFavorite(c *gin.Context) {
	userID := c.GetUint("user_id")

	bookIDParam := c.Param("id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	var book models.Book
	if err := config.DB.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	fav := models.FavoriteBook{
		UserID: userID,
		BookID: uint(bookID),
	}

	err = config.DB.FirstOrCreate(&fav, fav).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "added to favorites",
	})
}

func DeleteFavorite(c *gin.Context) {
	userID := c.GetUint("user_id")
	bookIDParam := c.Param("id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}
	result := config.DB.
		Where("user_id = ? AND book_id = ?", userID, bookID).
		Delete(&models.FavoriteBook{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "favorite not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
