package db

import (
	"log"

	"gorm.io/gorm"
)

func Cleanup(db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err == nil {
			log.Println("[closing database connections...]")
			if closeErr := sqlDB.Close(); closeErr != nil {
				log.Printf("[error closing database: %v]", closeErr)
			} else {
				log.Println("[database connections closed]")
			}
		}
	}
}
