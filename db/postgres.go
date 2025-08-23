package db

import (
	"app/server/config"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDBSource(config *config.Postgres) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.Database)
}

func InitPostgres(config *config.Postgres) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", config.Host, config.User, config.Password, config.Database, config.Port, "UTC")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("[postgres database connected]")
	return db
}

func RunMigration(config *config.Postgres) {
	dbSrc := getDBSource(config)
	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(currentFile)) // Go up two levels from db/run_migrations.go to project root
	migrationPath := filepath.Join(projectRoot, "migrations")
	// migrations are in root of the project one folder level up
	m, err := migrate.New(
		"file://"+migrationPath,
		dbSrc)

	if err != nil {
		log.Fatal("[failed to create migration instance]:", err)
	}
	if err := m.Up(); err != nil {
		log.Println("[failed to run migrations]:", err)
	}
	log.Println("[migrations script executed successfully]")
}
