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

func (ar *AuthRepo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	result := ar.db.WithContext(ctx).Create(&user)
	return user, result.Error
}

func (ar *AuthRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	result := ar.db.WithContext(ctx).Model(model.User{}).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
