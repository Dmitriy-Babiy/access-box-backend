package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// AuthMiddleware предоставляет middleware для аутентификации
type AuthMiddleware struct {
	// TODO: Добавить конфигурацию JWT
}

// NewAuthMiddleware создает новый middleware для аутентификации
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// AuthMiddleware возвращает gin.HandlerFunc для проверки аутентификации
func (m *AuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Реализовать проверку JWT токена
		token := extractToken(c)
		
		if token == "" {
			c.JSON(401, gin.H{"error": "Токен не предоставлен"})
			c.Abort()
			return
		}
		
		// TODO: Валидировать токен
		if !m.isValidToken(token) {
			c.JSON(401, gin.H{"error": "Недействительный токен"})
			c.Abort()
			return
		}
		
		// Токен валиден, продолжаем
		c.Next()
	}
}

// extractToken извлекает токен из заголовка Authorization
func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	
	// Формат: "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	
	return parts[1]
}

// isValidToken проверяет валидность токена
func (m *AuthMiddleware) isValidToken(token string) bool {
	// TODO: Реализовать проверку JWT токена
	// Пока возвращаем true для заглушки
	return true
}
