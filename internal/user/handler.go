package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler обрабатывает HTTP запросы для пользователей
type Handler struct {
	userService UserService
}

// NewHandler создает новый обработчик пользователей
func NewHandler(userService UserService) *Handler {
	return &Handler{userService: userService}
}

// GetUsers получает список пользователей
func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении пользователей"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// CreateUser создает нового пользователя
func (h *Handler) CreateUser(c *gin.Context) {
	var request struct {
		Email    string `json:"email" binding:"required,email"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}

	user, err := h.userService.CreateUser(request.Email, request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании пользователя"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// GetUser получает пользователя по ID
func (h *Handler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID пользователя"})
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser обновляет пользователя
func (h *Handler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID пользователя"})
		return
	}

	var request struct {
		Email    string `json:"email" binding:"required,email"`
		Username string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}

	user, err := h.userService.UpdateUser(uint(id), request.Email, request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// DeleteUser удаляет пользователя
func (h *Handler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID пользователя"})
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно удален"})
}
