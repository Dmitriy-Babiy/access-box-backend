package main

import (
	"log"

	"access-box/internal/shared/database"
	"access-box/internal/shared/middleware"
	"access-box/internal/shared/router"
	"access-box/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Ошибка инициализации БД:", err)
	}

	// Автомиграция моделей
	if err := db.InitSchema(); err != nil {
		log.Fatal("Ошибка автомиграции:", err)
	}

	// Инициализация репозиториев
	userRepo := user.NewUserRepository(db)

	// Инициализация сервисов
	userService := user.NewUserService(userRepo)

	// Инициализация middleware
	authMiddleware := middleware.NewAuthMiddleware()

	// Инициализация handlers
	userHandler := user.NewHandler(userService)

	// Настройка маршрутов
	router := router.SetupRouter(userHandler, authMiddleware)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	log.Println("Сервер запущен на :8080")
	log.Fatal(router.Run(":8080"))
}
