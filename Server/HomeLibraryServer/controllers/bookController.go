package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/Sandeshpd/home-library/Server/HomeLibraryServer/database"
	"github.com/Sandeshpd/home-library/Server/HomeLibraryServer/models"
)

// Store a book collection in a variable for global use

var validate = validator.New()

// Get a list of available Books
func GetBooks(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Timeout facility for resource release
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var bookCollection *mongo.Collection = database.OpenCollection("books", client)

		var books []models.Book
		cursor, err := bookCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books."})
			return
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &books); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode books."})
			return
		}

		c.JSON(http.StatusOK, books)
	}
}

// Get a single book associated with provided "book_id"
func GetBookById(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Timeout facility for resource release
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var bookCollection *mongo.Collection = database.OpenCollection("books", client)

		bookId := c.Param("book_id")

		if bookId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Book ID cannot be null"})
			return
		}
		var book models.Book

		err := bookCollection.FindOne(ctx, bson.M{"book_id": bookId}).Decode(&book)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found"})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

// Add a book in the database
func AddBook(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Timeout facility for resource release
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var book models.Book
		err := c.ShouldBindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := validate.Struct(book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation error", "details": err.Error()})
			return
		}

		var bookCollection *mongo.Collection = database.OpenCollection("books", client)

		result, err := bookCollection.InsertOne(ctx, book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Book created", "Result": result})
	}
}

// Update the book associated with specified ID
func UpdateBook(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Timeout facility for resource release
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var book models.Book

		err := c.ShouldBindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := validate.Struct(book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation error", "details": err.Error()})
			return
		}

		var bookCollection *mongo.Collection = database.OpenCollection("books", client)

		bookId := c.Param("book_id")
		if bookId == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found."})
			return
		}

		filter := bson.M{"book_id": bookId}

		update := bson.M{
			"$set": bson.M{
				"title":       book.Title,
				"author":      book.Author,
				"price":       book.Price,
				"cover_path":  book.CoverPath,
				"language":    book.Language,
				"publication": book.Publication,
				"genre":       book.Genre,
			},
		}

		result, err := bookCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit book."})
			return
		}

		if result.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book updation successful.", "Result": book})
	}
}

// Delete the book associated with specified ID
func DeleteBook(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Timeout facility for resource release
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		// var bookCollection *mongo.Collection = database.OpenCollection("books", client)

		bookId := c.Param("book_id")
		if bookId == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found."})
			return
		}

		var bookCollection *mongo.Collection = database.OpenCollection("books", client)

		filter := bson.M{"book_id": bookId}
		result, err := bookCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book."})
			return
		}

		if result.DeletedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book deletion successful.", "Result": result})
	}
}

// Get all genres from the database
func GetGenre(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var genreCollection *mongo.Collection = database.OpenCollection("genre", client)

		cursor, err := genreCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching genres."})
			return
		}
		defer cursor.Close(ctx)

		var genres []models.Genre
		if err := cursor.All(ctx, &genres); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, genres)
	}
}
