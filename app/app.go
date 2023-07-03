package app

import (
	"clothing-shop/app/config"
	"clothing-shop/handler"
	"clothing-shop/service"
)

type ApplicationContext struct {
	ProductHandler *handler.ProductHandler
	UserHandler    *handler.UserHandler
}

func NewApp() *ApplicationContext {
	db := config.ConnectToDB()
	productService := service.NewProductService(db)
	productHandler := handler.NewProductHandler(*productService)
	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(*userService)

	return &ApplicationContext{
		ProductHandler: productHandler,
		UserHandler:    userHandler,
	}
}
