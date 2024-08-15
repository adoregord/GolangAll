package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	SetRestRoutes(router)
	SetSoapRoutes(router)

	return router
}
