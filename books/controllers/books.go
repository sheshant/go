package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sheshant/go_programming_books/models"
	"gorm.io/gorm"
)

// Handlers for CRUD operations

// Create a new book
func createBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
			return
		}
		c.JSON(http.StatusCreated, book)
	}
}

// Get all books
// getBooks is a handler function that retrieves a list of books from the database
// and returns them as a JSON response. It uses the provided Gorm database instance
// to query the books table.
//
// Parameters:
//   - db: A pointer to a gorm.DB instance used to interact with the database.
//
// Returns:
//   - A gin.HandlerFunc that can be used as a route handler in a Gin application.
//
// Behavior:
//   - On success, it responds with an HTTP 200 status code and a JSON array of books.
//   - On failure, it responds with an HTTP 500 status code and a JSON error message.
func getBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Book
		if err := db.Find(&books).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}

// Get a single book by ID
func getBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var book models.Book
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

// Update a book
func updateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var book models.Book
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		var updatedBook models.Book
		if err := c.ShouldBindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		book.Title = updatedBook.Title
		book.Author = updatedBook.Author
		book.ISBN = updatedBook.ISBN
		if err := db.Save(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

// Delete a book
func deleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var book models.Book
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		if err := db.Delete(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
	}
}
