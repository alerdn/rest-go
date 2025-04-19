package user

type UserUsecase struct {
	repository UserRepository
}

func NewUserUsecase(repository UserRepository) UserUsecase {
	return UserUsecase{
		repository,
	}
}

func (uc *UserUsecase) GetUsers() ([]User, error) {
	return uc.repository.GetUsers()
}
