package route

import (
	"app/server/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

var (
	AuthRouteBasePath = "/auth"
)

func RegisterRoutes(rtr *gin.RouterGroup) {

	rtrGrp := rtr.Group(AuthRouteBasePath)
	{
		authController := handler.NewAuthController()
		rtrGrp.POST("/", authController.Register)
	}
}
