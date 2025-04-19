package product

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	ProductController := CreateProductController()

	g := api.Group("/products")
	{
		g.GET("/", ProductController.GetProduct)
		g.POST("/", ProductController.CreateProduct)
		g.GET("/:id", ProductController.GetProductById)
	}
}
