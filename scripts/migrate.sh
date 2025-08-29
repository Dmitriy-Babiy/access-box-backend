#!/bin/bash

# Скрипт для управления миграциями в development окружении

set -e

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Функция для вывода сообщений
log() {
    echo -e "${BLUE}[$(date +'%Y-%m-%d %H:%M:%S')]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# Проверяем, что .env файл существует
if [ ! -f .env ]; then
    error "Файл .env не найден. Создайте его на основе env.example"
    exit 1
fi

# Загружаем переменные окружения
export $(cat .env | grep -v '^#' | xargs)

# Проверяем, что Goose установлен
GOOSE_BIN=$(go env GOPATH)/bin/goose
if [ ! -f "$GOOSE_BIN" ]; then
    error "Goose не установлен. Установите его командой: make install-goose"
    exit 1
fi

# Директория миграций
MIGRATIONS_DIR="./migrations"

# Функция для показа справки
show_help() {
    echo "Использование: $0 [КОМАНДА]"
    echo ""
    echo "Команды:"
    echo "  up       - Применить все миграции"
    echo "  down     - Откатить последнюю миграцию"
    echo "  status   - Показать статус миграций"
    echo "  create   - Создать новую миграцию (требует NAME=имя_миграции)"
    echo "  reset    - Откатить все миграции (ОСТОРОЖНО!)"
    echo "  help     - Показать эту справку"
    echo ""
    echo "Примеры:"
    echo "  $0 up"
    echo "  $0 create NAME=add_new_table"
    echo "  $0 status"
}

# Функция для создания миграции
create_migration() {
    if [ -z "$NAME" ]; then
        error "Укажите имя миграции: NAME=имя_миграции"
        exit 1
    fi
    
    log "Создаем новую SQL миграцию: $NAME"
    $GOOSE_BIN -dir "$MIGRATIONS_DIR" create "$NAME" sql
    success "Миграция создана успешно"
}

# Функция для применения миграций
run_migrations() {
    log "Применяем миграции..."
    $GOOSE_BIN -dir "$MIGRATIONS_DIR" up
    success "Миграции применены успешно"
}

# Функция для отката миграции
rollback_migration() {
    log "Откатываем последнюю миграцию..."
    $GOOSE_BIN -dir "$MIGRATIONS_DIR" down
    success "Миграция откачена успешно"
}

# Функция для показа статуса
show_status() {
    log "Показываем статус миграций..."
    $GOOSE_BIN -dir "$MIGRATIONS_DIR" status
}

# Функция для сброса всех миграций
reset_migrations() {
    warning "ВНИМАНИЕ: Это откатит ВСЕ миграции!"
    read -p "Вы уверены? Введите 'yes' для подтверждения: " confirm
    
    if [ "$confirm" = "yes" ]; then
        log "Откатываем все миграции..."
        $GOOSE_BIN -dir "$MIGRATIONS_DIR" down-to 0
        success "Все миграции откачены"
    else
        log "Операция отменена"
    fi
}

# Основная логика
case "${1:-help}" in
    "up")
        run_migrations
        ;;
    "down")
        rollback_migration
        ;;
    "status")
        show_status
        ;;
    "create")
        create_migration
        ;;
    "reset")
        reset_migrations
        ;;
    "help"|*)
        show_help
        ;;
esac
