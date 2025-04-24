package factories

import (
	"github.com/alerdn/rest-go/config"
	"github.com/alerdn/rest-go/internal/controllers"
	"github.com/alerdn/rest-go/internal/repositories"
	"github.com/alerdn/rest-go/internal/services"
)

func CreateUserController() controllers.UserController {
	UserRepository := repositories.NewUserRepository(config.DB)
	UserUsecase := services.NewUserService(UserRepository)
	UserController := controllers.NewUserController(UserUsecase)

	return UserController
}
