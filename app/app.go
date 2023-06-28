package app

import (
	"clothing-shop/app/config"
	"clothing-shop/handler"
	"clothing-shop/service"
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
