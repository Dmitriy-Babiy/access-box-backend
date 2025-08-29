package router

import (
	"access-box/internal/shared/middleware"
	"access-box/internal/user"

	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает все маршруты приложения
func SetupRouter(userHandler *user.Handler, authMiddleware *middleware.AuthMiddleware) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Публичные маршруты
	public := router.Group("/api/v1")
	{
		// Маршруты пользователей
		users := public.Group("/users")
		{
			users.POST("/", userHandler.CreateUser)    // Создание пользователя
			users.GET("/", userHandler.GetUsers)       // Получение всех пользователей
			users.GET("/:id", userHandler.GetUser)     // Получение пользователя по ID
		}
	}

	// Защищенные маршруты (требуют аутентификации)
	protected := router.Group("/api/v1")
	protected.Use(authMiddleware.AuthMiddleware())
	{
		// Защищенные маршруты пользователей
		protectedUsers := protected.Group("/users")
		{
			protectedUsers.PUT("/:id", userHandler.UpdateUser)   // Обновление пользователя
			protectedUsers.DELETE("/:id", userHandler.DeleteUser) // Удаление пользователя
		}
	}

	return router
}
