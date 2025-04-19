package product

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	ProductController := CreateProductController()

	server.GET("/products", ProductController.GetProduct)
	server.POST("/products", ProductController.CreateProduct)
	server.GET("/products/:id", ProductController.GetProductById)
}
