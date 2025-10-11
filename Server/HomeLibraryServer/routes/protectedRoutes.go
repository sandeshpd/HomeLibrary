package routes

import (
	controller "github.com/Sandeshpd/home-library/Server/HomeLibraryServer/controllers"
	"github.com/Sandeshpd/home-library/Server/HomeLibraryServer/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	// Protect endpoints behind Token authorization
	router.Use(middleware.AuthMiddlware())

	router.GET("/book/:book_id", controller.GetBookById(client))
	router.POST("/book/add", controller.AddBook(client))
	router.PUT("/book/update/:book_id", controller.UpdateBook(client))
	
	// FIXME: Protection isn't working on this endpoint.
	router.DELETE("/book/delete/:book_id", controller.DeleteBook(client))
}
