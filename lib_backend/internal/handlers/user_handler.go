package handler

import (
	"log"
	"net/http"

	"lib_backend/internal/model"
	"lib_backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	createdUser, err := h.userService.CreateUser(&user)

	if err != nil {
		log.Printf("ERROR: CreateUser service failed: %v", err)

		if err.Error() == "service: user with email "+user.Email+" already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format", "details": err.Error()})
		return
	}

	user, err := h.userService.GetUserByID(id)

	if err != nil {
		log.Printf("ERROR: GetUserByID service failed for ID %s: %v", id.String(), err)

		if err.Error() == "service: user with ID "+id.String()+" not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email parameter is required"})
		return
	}

	user, err := h.userService.GetUserByEmail(email)

	if err != nil {
		log.Printf("ERROR: GetUserByEmail service failed for email %s: %v", email, err)

		if err.Error() == "service: user with email "+email+" not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format", "details": err.Error()})
		return
	}

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	user.ID = id

	updatedUser, err := h.userService.UpdateUser(&user)

	if err != nil {
		log.Printf("ERROR: UpdateUser service failed for ID %s: %v", id.String(), err)

		if err.Error() == "service: user with ID "+id.String()+" not found for update" { // Replace with proper error type checking
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format", "details": err.Error()})
		return
	}

	err = h.userService.DeleteUser(id)

	if err != nil {
		log.Printf("ERROR: DeleteUser service failed for ID %s: %v", id.String(), err)

		if err.Error() == "service: user with ID "+id.String()+" not found for update" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user", "details": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		log.Printf("ERROR: GetAllUsers service failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
