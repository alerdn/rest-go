package sale

type SaleUsecase struct {
	repository SaleRepository
}

func NewSaleUsecase(repository SaleRepository) SaleUsecase {
	return SaleUsecase{
		repository,
	}
}

func (su *SaleUsecase) GetSalesByUser(userId int) ([]Sale, error) {
	return su.repository.GetSalesByUser(userId)
}
