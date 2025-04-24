package repositories

import (
	"github.com/alerdn/rest-go/internal/models"
	"gorm.io/gorm"
)

type SaleRepository struct {
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) SaleRepository {
	return SaleRepository{
		db,
	}
}

func (sr *SaleRepository) Create(sale models.Sale) (models.Sale, error) {
	result := sr.db.Create(&sale)
	if result.Error != nil {
		return models.Sale{}, result.Error
	}

	return sale, nil
}

func (sr *SaleRepository) GetSalesByUser(userId int) ([]models.Sale, error) {
	var sales []models.Sale

	result := sr.db.Preload("Product").Find(&sales, "sale.user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}

	return sales, nil
}

func (sr *SaleRepository) GetNextNsu() int {
	var lastSale models.Sale
	result := sr.db.Order("nsu desc").First(&lastSale)
	if result.Error != nil {
		return 1
	}
	return lastSale.Nsu + 1
}
