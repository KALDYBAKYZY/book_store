package handlers

import (
	"ginExample/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var books = []models.Book{
	{ID: 1, Title: "Go Programming", AuthorID: 1, CategoryID: 1, Price: 29.99},
	{ID: 2, Title: "Go Programming2", AuthorID: 1, CategoryID: 1, Price: 39.99},
	{ID: 3, Title: "Go Programming3", AuthorID: 1, CategoryID: 1, Price: 49.99},
	{ID: 4, Title: "Go Programming4", AuthorID: 1, CategoryID: 1, Price: 59.99},
	{ID: 5, Title: "Go Programming5", AuthorID: 1, CategoryID: 1, Price: 69.99},
	{ID: 6, Title: "Go Programming6", AuthorID: 1, CategoryID: 1, Price: 79.99},
	{ID: 7, Title: "Go Programming7", AuthorID: 1, CategoryID: 1, Price: 89.99},
}

func GetBooks(c *gin.Context) {
	categoryFilter := c.Query("category")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	//limit := 5
	start := (page - 1) * limit
	end := start + limit

	var filteredBooks []models.Book
	for _, book := range books {
		if categoryFilter == "" || strconv.Itoa(book.CategoryID) == categoryFilter {
			filteredBooks = append(filteredBooks, book)
		}
	}

	if start >= len(filteredBooks) {
		c.JSON(http.StatusOK, []models.Book{})
		return
	}

	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	c.JSON(http.StatusOK, filteredBooks[start:end])
}

func AddBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newBook.ID = len(books) + 1
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedBook models.Book

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
