.PHONY: help build run test clean migrate-create migrate-up migrate-down migrate-status migrate-reset

# Переменные
BINARY_NAME=access-box
MIGRATIONS_DIR=./migrations
GOOSE_BIN=$(shell go env GOPATH)/bin/goose
AIR_BIN=$(shell go env GOPATH)/bin/air

# Помощь
help: ## Показать справку по командам
	@echo "Доступные команды:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Сборка
build: ## Собрать приложение
	go build -o $(BINARY_NAME) ./cmd/api

# Запуск
run: build ## Собрать и запустить приложение
	./$(BINARY_NAME)

# Live reload с Air
dev: ## Запустить приложение с live-reload (Air)
	@if [ ! -f "$(AIR_BIN)" ]; then \
		echo "❌ Air не установлен. Устанавливаю..."; \
		go install github.com/air-verse/air@latest; \
	fi
	$(AIR_BIN)

# Установка Air
install-air: ## Установить Air для live-reload
	go install github.com/air-verse/air@latest

# Проверка версии Air
air-version: ## Показать версию Air
	@echo "Air установлен в: $(AIR_BIN)"
	@echo "Air доступен: $(shell test -f $(AIR_BIN) && echo "✅ Да" || echo "❌ Нет")"

# Настройка логирования БД
db-log-silent: ## Отключить логи БД (silent)
	@echo "🔇 Отключаю логи БД..."
	@export DB_LOG_LEVEL=silent && echo "✅ DB_LOG_LEVEL установлен в 'silent'"

db-log-error: ## Включить логи БД только для ошибок
	@echo "⚠️  Включаю логи БД для ошибок..."
	@export DB_LOG_LEVEL=error && echo "✅ DB_LOG_LEVEL установлен в 'error'"

db-log-warn: ## Включить логи БД для предупреждений и ошибок
	@echo "⚠️  Включаю логи БД для предупреждений и ошибок..."
	@export DB_LOG_LEVEL=warn && echo "✅ DB_LOG_LEVEL установлен в 'warn'"

db-log-info: ## Включить все логи БД (SQL запросы)
	@echo "📝 Включаю все логи БД..."
	@export DB_LOG_LEVEL=info && echo "✅ DB_LOG_LEVEL установлен в 'info'"

db-log-status: ## Показать текущий уровень логирования БД
	@echo "🔍 Текущий уровень логирования БД:"
	@echo "DB_LOG_LEVEL: $(shell echo $$DB_LOG_LEVEL || echo 'не установлен (по умолчанию: silent)')"

# Тесты
test: ## Запустить тесты
	go test ./...

# Очистка
clean: ## Очистить собранные файлы
	rm -f $(BINARY_NAME)

# Миграции
migrate-create: ## Создать новую миграцию (использовать: make migrate-create NAME=migration_name)
	@if [ -z "$(NAME)" ]; then echo "Ошибка: укажите NAME=migration_name"; exit 1; fi
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) create $(NAME) sql

migrate-up: ## Применить все миграции
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) up

migrate-down: ## Откатить последнюю миграцию
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) down

migrate-status: ## Показать статус миграций
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) status

migrate-reset: ## Откатить все миграции (ОСТОРОЖНО!)
	@echo "⚠️  ВНИМАНИЕ: Это откатит ВСЕ миграции!"
	@read -p "Вы уверены? Введите 'yes' для подтверждения: " confirm; \
	if [ "$$confirm" = "yes" ]; then \
		$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) down-to 0; \
	else \
		echo "Операция отменена"; \
	fi

# Установка Goose
install-goose: ## Установить Goose CLI
	go install github.com/pressly/goose/v3/cmd/goose@latest

# Проверка версии Goose
goose-version: ## Показать версию Goose
	@echo "Goose установлен в: $(GOOSE_BIN)"
	@echo "Goose доступен: $(shell test -f $(GOOSE_BIN) && echo "✅ Да" || echo "❌ Нет")"

# Создание Go миграции
migrate-create-go: ## Создать новую Go миграцию (использовать: make migrate-create-go NAME=migration_name)
	@if [ -z "$(NAME)" ]; then echo "Ошибка: укажите NAME=migration_name"; exit 1; fi
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) create $(NAME) go

# Использование скрипта миграций
migrate-script: ## Запустить скрипт миграций (использовать: make migrate-script CMD=up|down|status|create|reset)
	@if [ -z "$(CMD)" ]; then echo "Ошибка: укажите CMD=команда"; exit 1; fi
	./scripts/migrate.sh $(CMD)

# Проверка структуры миграций
migrate-list: ## Показать список файлов миграций
	@echo "📁 Файлы миграций в директории $(MIGRATIONS_DIR):"
	@ls -la $(MIGRATIONS_DIR) || echo "Директория миграций не найдена"
