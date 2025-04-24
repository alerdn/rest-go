package services

import (
	"github.com/alerdn/rest-go/internal/models"
	"github.com/alerdn/rest-go/internal/repositories"
)

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return UserService{
		repository,
	}
}

func (uc *UserService) Create(user models.User) (models.User, error) {
	id, err := uc.repository.Create(user)
	if err != nil {
		return models.User{}, err
	}

	user.ID = id

	return user, nil
}

func (uc *UserService) GetUsers() ([]models.User, error) {
	return uc.repository.GetUsers()
}

func (uc *UserService) GetUserByEmail(email string) (models.User, error) {
	return uc.repository.GetUserByEmail(email)
}

func (uc *UserService) GetUserByID(id int) (models.User, error) {
	return uc.repository.GetUserByID(id)
}
