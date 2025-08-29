#!/bin/bash

echo "🚀 Запуск Access Box..."

# Проверяем, что PostgreSQL запущен
if ! pg_isready -h localhost -p 5432 -U postgres > /dev/null 2>&1; then
    echo "❌ PostgreSQL не доступен на localhost:5432"
    echo "Убедитесь, что PostgreSQL запущен и доступен"
    exit 1
fi

echo "✅ PostgreSQL доступен"

# Устанавливаем переменные окружения
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=access-box
export SERVER_PORT=8080

echo "🔧 Переменные окружения установлены:"
echo "   DB_HOST: $DB_HOST"
echo "   DB_PORT: $DB_PORT"
echo "   DB_USER: $DB_USER"
echo "   DB_NAME: $DB_NAME"
echo "   SERVER_PORT: $SERVER_PORT"

echo ""
echo "🌐 Приложение будет доступно по адресу: http://localhost:$SERVER_PORT"
echo ""

# Запускаем приложение
go run main.go
