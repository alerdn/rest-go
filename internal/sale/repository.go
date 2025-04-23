package sale

import "gorm.io/gorm"

type SaleRepository struct {
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) SaleRepository {
	return SaleRepository{
		db,
	}
}

func (sr *SaleRepository) GetSalesByUser(userId int) ([]Sale, error) {
	var sales []Sale

	result := sr.db.Find(&sales, "sale.user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}

	return sales, nil
}
