package main

import (
	"log"

	"access-box/internal/shared/database"
	"access-box/internal/shared/middleware"
	"access-box/internal/shared/router"
	"access-box/internal/user"
)

func main() {
	// Инициализация базы данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("❌ ошибка инициализации БД:", err)
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

	log.Println("🚀 Сервер готов к запуску на http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
