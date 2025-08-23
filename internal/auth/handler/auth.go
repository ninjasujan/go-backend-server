package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac *AuthController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "register route called")
}
