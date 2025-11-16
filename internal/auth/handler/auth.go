package handler

import (
	"app/server/internal/auth/data"
	authService "app/server/internal/auth/service"
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
	var req data.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ac.authService.Register(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, response)
	}
}
