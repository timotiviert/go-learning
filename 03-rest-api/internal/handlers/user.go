package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/timotiviert/go-learning/03-rest-api/internal/models"
	"github.com/timotiviert/go-learning/03-rest-api/internal/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var ru models.RegisterUsers
	if err := c.ShouldBindJSON(&ru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Create(&ru)
	if err != nil {
		// Here or rather in the service we should check what the error was, if i.e. username or email already exists.
		// Only when error cant be explained -> HTTP 500
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.service.GetByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": })
		return
	}
}
