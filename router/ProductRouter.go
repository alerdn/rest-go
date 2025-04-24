package router

import (
	"github.com/alerdn/rest-go/internal/factories"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(api *gin.RouterGroup) {
	ProductController := factories.CreateProductController()

	g := api.Group("/products")
	{
		g.GET("/", ProductController.GetProduct)
		g.POST("/", ProductController.CreateProduct)
		g.GET("/:id", ProductController.GetProductById)
	}
}
