package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Sandeshpd/home-library/Server/HomeLibraryServer/database"
	"github.com/Sandeshpd/home-library/Server/HomeLibraryServer/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	router := gin.Default()

	router.GET("/intro", func(c *gin.Context) {
		c.String(200, "Hello User from Home Library Server...")
	})

	// TODO: Register CORS modules here
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: Unable to find .env file.")
	}

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	var origins []string
	if allowedOrigins != "" {
		origins = strings.Split(allowedOrigins, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
			log.Println("Allowed origin:", origins[i])
		}
	} else {
		origins = []string{"http://localhost:5173"}
		log.Println("Allowed origin: http://localhost:5173")
	}

	config := cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))
	router.Use(gin.Logger())

	// Establish database connection
	var client *mongo.Client = database.Connect()
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to reach the server: %v", err)
	}
	// Disconnect the database after making connection. Defer means 
	// it delays the execution until nearby methods execute.
	defer func() {
		err:=client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Failed to disconnect from database: %v", err)
		}
	}()

	routes.SetupUnprotectedRoutes(router, client)
	routes.SetupProtectedRoutes(router, client)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start the server", err)
	}

}
