package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/erpachecomo/photo-api/internal/models"
	appErrors "github.com/erpachecomo/photo-api/pkg/utils"
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
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) PostUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("BindJSON Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	newUser, err := h.svc.PostUser(&user)

	if errors.Is(err, appErrors.ErrEntityAlreadyExists) {
		c.JSON(http.StatusConflict, gin.H{"message": "Email already exists"})
		return
	}

	if err != nil {
		fmt.Printf("Post Service Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
func (h *UserHandler) PutUser(c *gin.Context) {
	var user models.UserUpdate
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("BindJSON Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	id := c.Param("id")
	updatedUser, err := h.svc.PutUser(id, &user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	if updatedUser == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No changes made"})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.svc.DeleteUser(id)
	if err != nil {
		fmt.Printf("Delete Service Error: %v", err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
func (h *UserHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}
