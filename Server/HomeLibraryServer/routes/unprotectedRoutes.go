package routes

import (
	controller "github.com/Sandeshpd/home-library/Server/HomeLibraryServer/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupUnprotectedRoutes(router *gin.Engine, client *mongo.Client)  {
	router.GET("/books", controller.GetBooks(client))
	router.GET("/books/genres", controller.GetGenre(client))

	router.POST("/user/register", controller.RegisterUser(client))
	router.POST("/user/login", controller.LoginUser(client))
	router.POST("/user/logout", controller.LogoutHandler(client))
}