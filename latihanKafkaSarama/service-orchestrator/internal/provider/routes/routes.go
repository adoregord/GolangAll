package routes

import (
	"kafka-sarama-try/internal/handler"
	"kafka-sarama-try/internal/orchestrator"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	orderRoutes := router.Group("/order")
	{
		orderRoutes.POST("/", handler.OrderHandler)
	}

	go func() {
		orchestrator.StartConsumer()
	}()

	return router
}
