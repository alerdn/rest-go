package repositories

import (
	"log"

	"github.com/alerdn/rest-go/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db,
	}
}

func (ur *UserRepository) Create(user models.User) (int, error) {
	result := ur.db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return 0, result.Error
	}

	return user.ID, nil
}

func (ur *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User

	result := ur.db.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	result := ur.db.Find(&user, "email = ?", email)
	if result.Error != nil {
		log.Println(result.Error)
		return models.User{}, result.Error
	}

	return user, nil
}

func (ur *UserRepository) GetUserByID(id int) (models.User, error) {
	var user models.User

	result := ur.db.Preload("Sales").Find(&user, id)
	if result.Error != nil {
		log.Println(result.Error)
		return models.User{}, result.Error
	}

	return user, nil
}
