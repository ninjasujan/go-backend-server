package repo

import (
	"context"

	"app/server/internal/auth/model"

	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db}
}

func (ar *AuthRepo) CreateUser(ctx context.Context, user *model.Auth) error {
	return ar.db.Create(user).Error
}
