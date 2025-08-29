# Управление миграциями базы данных

Этот проект использует [Goose](https://github.com/pressly/goose) для управления миграциями базы данных.

## Установка Goose

```bash
# Установка через Go
go install github.com/pressly/goose/v3/cmd/goose@latest

# Или через Makefile
make install-goose
```

## Структура миграций

Миграции находятся в директории `./migrations/` и могут быть написаны на SQL или Go.

### SQL миграции

SQL миграции должны содержать специальные комментарии:

```sql
-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;
```

### Go миграции

Go миграции должны содержать функции `Up` и `Down`:

```go
package migrations

import (
    "database/sql"
    "github.com/pressly/goose/v3"
)

func init() {
    goose.AddMigration(Up, Down)
}

func Up(tx *sql.Tx) error {
    // Логика миграции
    return nil
}

func Down(tx *sql.Tx) error {
    // Логика отката
    return nil
}
```

## Команды управления миграциями

### Через Makefile (рекомендуется)

```bash
# Показать справку
make help

# Создать новую SQL миграцию
make migrate-create NAME=add_new_table

# Создать новую Go миграцию
make migrate-create-go NAME=seed_data

# Применить все миграции
make migrate-up

# Откатить последнюю миграцию
make migrate-down

# Показать статус миграций
make migrate-status

# Откатить все миграции (ОСТОРОЖНО!)
make migrate-reset
```

### Через Goose CLI

```bash
# Создать миграцию
goose -dir ./migrations create migration_name sql
goose -dir ./migrations create migration_name go

# Применить миграции
goose -dir ./migrations up

# Откатить миграцию
goose -dir ./migrations down

# Показать статус
goose -dir ./migrations status

# Применить до определенной версии
goose -dir ./migrations up-to 20231201120000

# Откатить до определенной версии
goose -dir ./migrations down-to 20231201120000
```

## Переменные окружения

Создайте файл `.env` на основе `env.example`:

```bash
cp env.example .env
```

Настройте переменные в `.env`:

```env
# Настройки базы данных
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=access-box

# Настройки Goose
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://postgres:postgres@localhost:5432/access-box?sslmode=disable
GOOSE_MIGRATION_DIR=./migrations
GOOSE_TABLE=goose_migrations
```

## Автоматический запуск миграций

При запуске приложения миграции запускаются автоматически через `database.InitDB()`.

## Лучшие практики

1. **Именование файлов**: Используйте временные метки для версионирования
2. **Обратная совместимость**: Всегда пишите функции `Down` для отката
3. **Транзакции**: По умолчанию миграции выполняются в транзакциях
4. **Тестирование**: Тестируйте миграции на тестовой базе данных
5. **Резервные копии**: Делайте резервные копии перед применением миграций в продакшене

## Решение проблем

### Ошибка "driver: bad connection"

Проверьте настройки подключения к базе данных в `.env` файле.

### Ошибка "migration file not found"

Убедитесь, что директория миграций указана правильно: `GOOSE_MIGRATION_DIR=./migrations`

### Ошибка "duplicate migration version"

Проверьте, что все файлы миграций имеют уникальные версии.

## Полезные ссылки

- [Документация Goose](https://github.com/pressly/goose)
- [Примеры миграций](https://github.com/pressly/goose/tree/main/examples)
- [Поддерживаемые диалекты SQL](https://github.com/pressly/goose/blob/main/dialect.go)
