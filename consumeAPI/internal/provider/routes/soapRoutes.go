package routes

import (
	"consume-api-try/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetSoapRoutes(router *gin.Engine) {
	soapHandler := handler.NewSoapHandler()

	soapRoutes := router.Group("/soap")
	{
		soapRoutes.POST("/", soapHandler.ConsumeSOAP)
	}
}
