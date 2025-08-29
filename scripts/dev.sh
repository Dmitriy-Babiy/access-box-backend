#!/bin/bash

# Скрипт для запуска разработки с Air
# Автоматически устанавливает Air если не установлен

set -e

echo "🚀 Запуск режима разработки с Air..."

# Проверяем, установлен ли Air
if ! command -v air &> /dev/null; then
    echo "📦 Air не установлен. Устанавливаю..."
    go install github.com/air-verse/air@latest
    echo "✅ Air установлен успешно!"
else
    echo "✅ Air уже установлен"
fi

# Проверяем наличие конфигурационного файла
if [ ! -f ".air.toml" ]; then
    echo "❌ Файл .air.toml не найден!"
    echo "Создайте конфигурацию Air или используйте 'air init'"
    exit 1
fi

echo "🔧 Запускаю Air с конфигурацией .air.toml..."
echo "📁 Отслеживаемые расширения: .go, .tpl, .tmpl, .html"
echo "🚫 Исключенные директории: assets, tmp, vendor, testdata, migrations, .git"
echo ""

# Запускаем Air
air
