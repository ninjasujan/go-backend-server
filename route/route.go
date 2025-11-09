package route

import (
	authRoute "app/server/internal/auth/route"

	"app/server/common/kafka/producer"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB, engine *gin.Engine, kafkaProducer *producer.KafkaProducer) {

	api := engine.Group("/api")
	v1 := api.Group("/v1")

	authRoute.RegisterRoutes(db, v1, kafkaProducer)

}
