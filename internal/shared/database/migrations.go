package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pressly/goose/v3"
)

// MigrationManager управляет миграциями базы данных
type MigrationManager struct {
	db *sql.DB
}

// NewMigrationManager создает новый менеджер миграций
func NewMigrationManager(db *sql.DB) *MigrationManager {
	return &MigrationManager{db: db}
}

// RunMigrations запускает все доступные миграции
func (m *MigrationManager) RunMigrations(migrationsDir string) error {
	// Устанавливаем диалект PostgreSQL
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}

	// Запускаем миграции
	if err := goose.Up(m.db, migrationsDir); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}

// GetMigrationStatus возвращает статус миграций
func (m *MigrationManager) GetMigrationStatus(migrationsDir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}

	if err := goose.Status(m.db, migrationsDir); err != nil {
		return fmt.Errorf("failed to get migration status: %w", err)
	}

	return nil
}

// RollbackLastMigration откатывает последнюю миграцию
func (m *MigrationManager) RollbackLastMigration(migrationsDir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}

	if err := goose.Down(m.db, migrationsDir); err != nil {
		return fmt.Errorf("failed to rollback migration: %w", err)
	}

	log.Println("Last migration rolled back successfully")
	return nil
}

// GetCurrentVersion возвращает текущую версию миграций
func (m *MigrationManager) GetCurrentVersion() (int64, error) {
	if err := goose.SetDialect("postgres"); err != nil {
		return 0, fmt.Errorf("failed to set dialect: %w", err)
	}

	version, err := goose.GetDBVersion(m.db)
	if err != nil {
		return 0, fmt.Errorf("failed to get current version: %w", err)
	}

	return version, nil
}
