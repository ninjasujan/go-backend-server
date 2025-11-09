package service

import (
	"app/server/common/constant"
	"app/server/common/kafka/producer"
	"app/server/internal/auth/repo"
)

type AuthService struct {
	repo          *repo.AuthRepo
	kafkaProducer producer.Producer
}

func NewAuthService(repo *repo.AuthRepo, kafkaProducer producer.Producer) *AuthService {
	return &AuthService{repo: repo, kafkaProducer: kafkaProducer}
}

func (as *AuthService) Register() error {
	key := "user_registration"
	err := as.kafkaProducer.Publish(constant.KafkaTopic, key, []byte("User registered"))
	if err != nil {
		return err
	}
	return nil
}
