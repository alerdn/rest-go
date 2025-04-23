package sale

import (
	"github.com/alerdn/rest-go/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	SaleController := CreateSaleController()

	g := api.Group("/sales")
	{
		protected := g.Use(auth.Middleware())
		{
			protected.GET("/", SaleController.GetSales)
		}
	}
}
