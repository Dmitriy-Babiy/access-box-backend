-- +goose Up
-- +goose StatementBegin
-- Создаем таблицу для отслеживания миграций Goose
-- Эта таблица нужна для работы системы миграций
CREATE TABLE IF NOT EXISTS goose_migrations (
    id SERIAL PRIMARY KEY,
    version_id BIGINT NOT NULL,
    is_applied BOOLEAN NOT NULL DEFAULT true,
    tstamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создаем уникальный индекс на version_id для предотвращения дублирования
CREATE UNIQUE INDEX IF NOT EXISTS idx_goose_migrations_version_id ON goose_migrations(version_id);

-- Добавляем комментарий к таблице
COMMENT ON TABLE goose_migrations IS 'Таблица для отслеживания примененных миграций Goose';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Откат: удаляем таблицу и индекс
DROP INDEX IF EXISTS idx_goose_migrations_version_id;
DROP TABLE IF EXISTS goose_migrations;
-- +goose StatementEnd
