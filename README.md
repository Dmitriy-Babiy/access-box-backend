# Access Box

Сервис управления доступом с использованием Go, Gin и PostgreSQL.

## Особенности

- 🚀 Быстрый и легкий веб-фреймворк Gin
- 🗄️ PostgreSQL с автоматическими миграциями
- 🔐 JWT аутентификация
- 🏗️ Чистая архитектура с разделением на слои
- 📝 Автоматическое управление миграциями через Goose

## Требования

- Go 1.23+
- PostgreSQL 12+
- Make (для удобства разработки)

## Установка

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd access-box
```

2. Установите зависимости:
```bash
go mod download
```

3. Установите Goose для управления миграциями:
```bash
make install-goose
```

4. Настройте переменные окружения:
```bash
cp env.example .env
# Отредактируйте .env файл под ваши настройки
```

## Запуск

### Разработка

```bash
# Собрать и запустить
make run

# Или по отдельности
make build
./access-box
```

### Продакшн

```bash
# Собрать
make build

# Запустить
./access-box
```

## Управление миграциями

Этот проект использует [Goose](https://github.com/pressly/goose) для управления миграциями базы данных.

### Быстрый старт

```bash
# Показать справку по командам
make help

# Создать новую миграцию
make migrate-create NAME=add_new_feature

# Применить миграции
make migrate-up

# Показать статус
make migrate-status
```

### Через Makefile

```bash
# Создание миграций
make migrate-create NAME=add_users_table
make migrate-create-go NAME=seed_data

# Управление миграциями
make migrate-up          # Применить все
make migrate-down        # Откатить последнюю
make migrate-status      # Показать статус
make migrate-reset       # Откатить все (ОСТОРОЖНО!)

# Проверка
make migrate-list        # Список файлов
make goose-version       # Версия Goose
```

### Через скрипт

```bash
# Использование скрипта миграций
make migrate-script CMD=up
make migrate-script CMD=status
make migrate-script CMD=create NAME=test_migration

# Или напрямую
./scripts/migrate.sh up
./scripts/migrate.sh status
./scripts/migrate.sh create NAME=test_migration
```

### Через Goose CLI

```bash
# Настройка переменных окружения
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="postgres://user:password@localhost:5432/dbname?sslmode=disable"
export GOOSE_MIGRATION_DIR=./migrations

# Команды
goose up              # Применить все
goose down            # Откатить последнюю
goose status          # Показать статус
goose create init sql # Создать миграцию
```

## Структура проекта

```
access-box/
├── cmd/api/              # Точка входа приложения
├── internal/              # Внутренний код приложения
│   ├── shared/           # Общие компоненты
│   │   ├── database/     # Настройки БД и миграции
│   │   ├── middleware/   # HTTP middleware
│   │   └── router/       # Маршрутизация
│   └── user/             # Модуль пользователей
├── migrations/            # Миграции базы данных
├── pkg/                   # Публичные пакеты
├── scripts/               # Скрипты для разработки
├── Makefile               # Команды для разработки
├── env.example            # Пример переменных окружения
└── MIGRATIONS.md          # Документация по миграциям
```

## Переменные окружения

Создайте файл `.env` на основе `env.example`:

```env
# База данных
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=access-box

# Сервер
SERVER_PORT=8080

# Goose (опционально)
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://postgres:postgres@localhost:5432/access-box?sslmode=disable
GOOSE_MIGRATION_DIR=./migrations
GOOSE_TABLE=goose_migrations
```

## Разработка

### Создание новой миграции

```bash
# SQL миграция
make migrate-create NAME=add_new_table

# Go миграция
make migrate-create-go NAME=complex_logic
```

### Тестирование

```bash
make test
```

### Очистка

```bash
make clean
```

## Миграции

При запуске приложения миграции применяются автоматически. Подробная документация по миграциям находится в [MIGRATIONS.md](MIGRATIONS.md).

## Лицензия

MIT License

## Поддержка

Если у вас есть вопросы или проблемы, создайте issue в репозитории.
# access-box-backend
