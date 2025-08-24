package handler

import (
	"app/server/internal/auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (ac *AuthHandler) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "register route called")
}
