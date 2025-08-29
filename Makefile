.PHONY: help build run test clean migrate-create migrate-up migrate-down migrate-status migrate-reset

# Переменные
BINARY_NAME=access-box
MIGRATIONS_DIR=./migrations
GOOSE_BIN=$(shell go env GOPATH)/bin/goose

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
