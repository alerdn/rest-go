package router

import (
	"github.com/alerdn/rest-go/internal/product"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {

	server := gin.Default()

	product.RegisterRoutes(server)

	return server
}
