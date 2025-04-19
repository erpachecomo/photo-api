package api

import (
	"net/http"

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

	// Start the server
	router.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
