package user

type UserUsecase struct {
	repository UserRepository
}

func NewUserUsecase(repository UserRepository) UserUsecase {
	return UserUsecase{
		repository,
	}
}

func (uc *UserUsecase) Create(user User) (User, error) {
	id, err := uc.repository.Create(user)
	if err != nil {
		return User{}, err
	}

	user.ID = id

	return user, nil
}

func (uc *UserUsecase) GetUsers() ([]User, error) {
	return uc.repository.GetUsers()
}

func (uc *UserUsecase) GetUserByEmail(email string) (User, error) {
	return uc.repository.GetUserByEmail(email)
}

func (uc *UserUsecase) GetUserByID(id int) (User, error) {
	return uc.repository.GetUserByID(id)
}
