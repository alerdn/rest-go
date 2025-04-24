package router

import (
	"github.com/alerdn/rest-go/internal/factories"
	"github.com/alerdn/rest-go/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterSaleRoutes(api *gin.RouterGroup) {
	SaleController := factories.CreateSaleController()

	g := api.Group("/sales")
	{
		protected := g.Use(middlewares.Middleware())
		{
			protected.GET("/", SaleController.GetSales)
			protected.POST("/", SaleController.CreateSale)
		}
	}
}
