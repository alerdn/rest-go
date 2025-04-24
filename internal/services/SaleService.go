package services

import (
	"github.com/alerdn/rest-go/internal/models"
	"github.com/alerdn/rest-go/internal/repositories"
)

type SaleService struct {
	repository        repositories.SaleRepository
	productRepository repositories.ProductRepository
}

type CreateRequest struct {
	ProductId     int
	Quantity      int
	PaymentMethod string
	UserID        int
}

func NewSaleService(repository repositories.SaleRepository, productRepository repositories.ProductRepository) SaleService {
	return SaleService{
		repository,
		productRepository,
	}
}

func (su SaleService) Create(req CreateRequest) (models.Sale, error) {
	product, _ := su.productRepository.GetProductById(req.ProductId)

	sale := models.Sale{
		Nsu:           su.repository.GetNextNsu(),
		ProductId:     req.ProductId,
		UserId:        req.UserID,
		Price:         product.Price * float64(req.Quantity),
		PaymentMethod: req.PaymentMethod,
		Quantity:      req.Quantity,
	}

	newSale, err := su.repository.Create(sale)
	if err != nil {
		return models.Sale{}, err
	}

	return newSale, nil
}

func (su *SaleService) GetSalesByUser(userId int) ([]models.Sale, error) {
	return su.repository.GetSalesByUser(userId)
}
