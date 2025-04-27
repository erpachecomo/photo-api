package api

import (
	"net/http"
	"os"

	"github.com/erpachecomo/photo-api/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupRoutes(db *mongo.Database) {
	userHandler := handlers.NewUserHandler(db)
	// Create a new Gin router
	router := gin.Default()

	// Define the routes
	router.GET("/ping", ping)
	router.GET("/users/:id", userHandler.GetUser)
	router.POST("/users", userHandler.PostUser)
	router.PUT("/users/:id", userHandler.PutUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
	router.GET("/users", userHandler.GetUsers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}
	// Start the server
	router.Run("0.0.0.0:" + port)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
