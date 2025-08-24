package service

import "app/server/internal/auth/repo"

type AuthService struct {
	repo *repo.AuthRepo
}

func NewAuthService(repo *repo.AuthRepo) *AuthService {
	return &AuthService{repo: repo}
}
