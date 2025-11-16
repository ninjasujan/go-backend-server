package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID      `json:"id" gorm:"id primaryKey"`
	Email     string         `json:"email" validate:"required,email" gorm:"email unique"`
	FirstName string         `json:"firstName" validate:"required" gorm:"first_name not null"`
	LastName  string         `json:"lastName" validate:"required" gorm:"last_name not null"`
	Password  string         `json:"password" validate:"required,min=8" gorm:"password not null"`
	Address   string         `json:"address" gorm:"address not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"created_at not null"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"updated_at not null"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"deleted_at"`
	CreatedBy string         `json:"createdBy" gorm:"created_by"`
	UpdatedBy string         `json:"updatedBy" gorm:"updated_by"`
	DeletedBy string         `json:"deletedBy" gorm:"deleted_by"`
}
