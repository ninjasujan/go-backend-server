package route

import (
	"app/server/internal/auth/handler"
	repo2 "app/server/internal/auth/repo"
	service2 "app/server/internal/auth/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	AuthRouteBasePath = "/auth"
)

func RegisterRoutes(db *gorm.DB, rtr *gin.RouterGroup) {

	// get all necessary dependencies
	repo := repo2.NewAuthRepo(db)
	service := service2.NewAuthService(repo)

	rtrGrp := rtr.Group(AuthRouteBasePath)
	{
		authController := handler.NewAuthHandler(service)
		rtrGrp.POST("/", authController.Register)
	}
}
