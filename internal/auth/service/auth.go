package service

import (
	"app/server/common/kafka/producer"
	"app/server/internal/auth/data"
	"app/server/internal/auth/repo"
	"context"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, data *data.RegisterRequest) error
}

type AuthService struct {
	repo          *repo.AuthRepo
	kafkaProducer producer.Producer
}

func NewAuthService(repo *repo.AuthRepo, kafkaProducer producer.Producer) AuthServiceInterface {
	return &AuthService{repo: repo, kafkaProducer: kafkaProducer}
}

func (as *AuthService) Register(ctx context.Context, data *data.RegisterRequest) error {

	return nil

}
