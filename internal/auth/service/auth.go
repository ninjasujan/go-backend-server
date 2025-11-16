package service

import (
	"app/server/common/kafka/producer"
	"app/server/internal/auth/data"
	"app/server/internal/auth/model"
	"app/server/internal/auth/repo"
	"context"

	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, data *data.RegisterUserRequest) (*data.RegisterUserResponse, error)
}

type AuthService struct {
	repo          *repo.AuthRepo
	kafkaProducer producer.Producer
}

func NewAuthService(repo *repo.AuthRepo, kafkaProducer producer.Producer) AuthServiceInterface {
	return &AuthService{repo: repo, kafkaProducer: kafkaProducer}
}

func (as *AuthService) Register(ctx context.Context, req *data.RegisterUserRequest) (*data.RegisterUserResponse, error) {

	user, err := as.repo.GetUserByEmail(ctx, req.Email)
	if err == nil && user != nil {
		return nil, errors.New("user already exists")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Id:        uuid.New(),
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  string(hashedPassword),
		Address:   req.Address,
	}

	// create user
	createdUser, err := as.repo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &data.RegisterUserResponse{
		Status:  "success",
		Message: "User created successfully",
		Data:    createdUser,
	}, nil
}
