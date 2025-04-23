package sale

import "github.com/alerdn/rest-go/config"

func CreateSaleController() SaleController {
	saleRepository := NewSaleRepository(config.DB)
	saleUsecase := NewSaleUsecase(saleRepository)
	saleController := NewSaleController(saleUsecase)

	return saleController
}
