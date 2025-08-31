package db

import (
	"app/server/common/config"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDBSource(cfg config.Postgres) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}

func InitPostgres(cfg config.Postgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port, "UTC")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set connection pool parameters
	sqlDB.SetMaxIdleConns(10)           // Maximum idle connections
	sqlDB.SetMaxOpenConns(100)          // Maximum open connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Maximum connection lifetime

	log.Println("[postgres database connected with connection pool configured]")
	return db, nil
}

func RunMigration(cfg config.Postgres) error {
	dbSrc := getDBSource(cfg)
	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(currentFile)) // Go up two levels from db/run_migrations.go to project root
	migrationPath := filepath.Join(projectRoot, "migrations")
	// migrations are in root of the project one folder level up
	m, err := migrate.New(
		"file://"+migrationPath,
		dbSrc)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("[migrations are up to date]")
		} else {
			return fmt.Errorf("migration failed: %w", err)
		}
	} else {
		log.Println("[migrations executed successfully]")
	}
	return nil
}
