package handlers

import (
	"fmt"
	"net/http"

	"github.com/erpachecomo/photo-api/internal/models"
	"github.com/erpachecomo/photo-api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserHandler struct {
	svc *services.UserService
}

func NewUserHandler(db *mongo.Database) *UserHandler {
	return &UserHandler{
		svc: services.NewUserService(db),
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.svc.GetUser(id)
	fmt.Println(user)
	fmt.Println(err)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		panic(err)
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) PostUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("Binding Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newUser, err := h.svc.PostUser(&user)

	if err != nil {
		fmt.Printf("Post Service Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
