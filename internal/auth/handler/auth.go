package handler

import (
	"app/server/internal/auth/data"
	authService "app/server/internal/auth/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService authService.AuthServiceInterface
}

func NewAuthHandler(service authService.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (ac *AuthHandler) Register(ctx *gin.Context) {
	var req data.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.authService.Register(context.Background(), &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "register route called"})
}
