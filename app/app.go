package app

import (
	"clothing-shop/app/config"
	"clothing-shop/handler"
	"clothing-shop/service"
)

type ApplicationContext struct {
	ProductHandler *handler.ProductHandler
	UserHandler    *handler.UserHandler
	PartnerHandler *handler.PartnerHandler
}

func NewApp() *ApplicationContext {
	db := config.ConnectToDB()
	productService := service.NewProductService(db)
	productHandler := handler.NewProductHandler(*productService)
	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(*userService)
	partnerService := service.NewPartnerService(db)
	partnerHandler := handler.NewPartnerHandler(*partnerService)
	return &ApplicationContext{
		ProductHandler: productHandler,
		UserHandler:    userHandler,
		PartnerHandler : partnerHandler,
		
	}
}
