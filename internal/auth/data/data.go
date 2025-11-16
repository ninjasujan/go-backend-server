package data

import "app/server/internal/auth/model"

type RegisterUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Password  string `json:"password" binding:"required,min=8"`
	Address   string `json:"address" binding:"required"`
}

type RegisterUserResponse struct {
	Status  string      `json:"status" oneof:"success,error"`
	Message string      `json:"message"`
	Data    *model.User `json:"data"`
}
