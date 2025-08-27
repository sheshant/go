package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sheshant/go_programming_books/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initialize database
func setupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Auto-migrate the schema
	db.AutoMigrate(&models.Book{})
	return db
}

func main() {
	// Initialize Gin router
	r := gin.Default()
	// Initialize database
	db := setupDatabase()

	// Routes
	books := r.Group("/books")
	{
		books.POST("/", controllers.createBook(db))
		books.GET("/", controllers.getBooks(db))
		books.GET("/:id", controllers.getBook(db))
		books.PUT("/:id", controllers.updateBook(db))
		books.DELETE("/:id", controllers.deleteBook(db))
	}

	// Start server
	r.Run(":8080")
}
