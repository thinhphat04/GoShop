package handlers

import (
	"API/internal/models"
	"API/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Add a new user to the database with role, hashed password, and validation
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User details"
// @Success      201   {object}  models.User
// @Failure      400   {object}  map[string]interface{}  "Invalid request payload"
// @Failure      500   {object}  map[string]interface{}  "Failed to process request"
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Hash the password
	err := user.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Nếu không có `role` trong request, đặt mặc định là "user"
	if user.Role == "" {
		user.Role = "user"
	}

	err = h.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// GetUsers godoc
// @Summary      Retrieve all users
// @Description  Get a list of all registered users from the database
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200   {array}   models.User
// @Failure      500   {object}  map[string]interface{}  "Failed to retrieve users"
// @Router       /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
