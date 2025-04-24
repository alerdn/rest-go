package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {

	server := gin.Default()

	api := server.Group("/api")
	{
		RegisterProductRoutes(api)
		RegisterUserRoutes(api)
		RegisterSaleRoutes(api)
	}

	return server
}
