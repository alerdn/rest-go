package user

import "github.com/alerdn/rest-go/config"

func CreateUserController() UserController {
	UserRepository := NewUserRepository(config.DB)
	UserUsecase := NewUserUsecase(UserRepository)
	UserController := NewUserController(UserUsecase)

	return UserController
}
