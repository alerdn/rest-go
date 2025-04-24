package factories

import (
	"github.com/alerdn/rest-go/config"
	"github.com/alerdn/rest-go/internal/controllers"
	"github.com/alerdn/rest-go/internal/repositories"
	"github.com/alerdn/rest-go/internal/services"
)

func CreateSaleController() controllers.SaleController {
	saleRepository := repositories.NewSaleRepository(config.DB)
	productRepository := repositories.NewProductRepository(config.DB)
	saleUsecase := services.NewSaleService(saleRepository, productRepository)
	saleController := controllers.NewSaleController(saleUsecase)

	return saleController
}
