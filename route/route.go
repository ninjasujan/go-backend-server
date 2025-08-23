package route

import (
	"app/server/internal/auth/route"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {

	api := engine.Group("/api")
	v1 := api.Group("/v1")

	route.RegisterRoutes(v1)

}
