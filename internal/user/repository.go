package user

import (
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return UserRepository{
		connection,
	}
}

func (ur *UserRepository) Create(user User) (int, error) {
	result := ur.connection.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return 0, result.Error
	}

	return user.ID, nil
}

func (ur *UserRepository) GetUsers() ([]User, error) {
	var users []User

	result := ur.connection.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (User, error) {
	var user User

	result := ur.connection.Find(&user, "email = ?", email)
	if result.Error != nil {
		log.Println(result.Error)
		return User{}, result.Error
	}

	return user, nil
}

func (ur *UserRepository) GetUserByID(id int) (User, error) {
	var user User

	result := ur.connection.Find(&user, id)
	if result.Error != nil {
		log.Println(result.Error)
		return User{}, result.Error
	}

	return user, nil
}
