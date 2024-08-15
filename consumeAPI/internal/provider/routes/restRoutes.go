package routes

import (
	"consume-api-try/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetRestRoutes(router *gin.Engine) {
	restHandler := handler.NewRestHandler()

	restRoutes := router.Group("/rest")
	{
		restRoutes.GET("/", restHandler.ConsumeRest)
	}
}
