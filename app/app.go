package app

import (
	"onboarding/exercise1/app/config"
	"onboarding/exercise1/handler"
	"onboarding/exercise1/service"
)

type ApplicationContext struct {
	UserHandler *handler.UserHandler
}

func NewApp() *ApplicationContext {
	db := config.ConnectToDB()
	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(*userService)

	return &ApplicationContext{UserHandler: userHandler}
}
