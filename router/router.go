package router

import (
	"github.com/alerdn/rest-go/internal/product"
	"github.com/alerdn/rest-go/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {

	server := gin.Default()

	api := server.Group("/api")
	{
		product.RegisterRoutes(api)
		user.RegisterRoutes(api)
	}

	return server
}
