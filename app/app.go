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
	CategoryParentHandler *handler.CategoryParentHandler
	CategoryChildHandler handler.CategoryChildHandler 
}

func NewApp() *ApplicationContext {
	db := config.ConnectToDB()
	productService := service.NewProductService(db)
	productHandler := handler.NewProductHandler(*productService)
	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(*userService)
	partnerService := service.NewPartnerService(db)
	partnerHandler := handler.NewPartnerHandler(*partnerService)
	categoryParentService := service.NewCategoryParentService(db)
	categoryParentHandler := handler.NewCategoryParentHandler(*categoryParentService)
	categoryChildService := service.NewCategoryChildService(db)
	categoryChildHandler := handler.NewCategoryChildHandler(*categoryChildService)
	return &ApplicationContext{
		ProductHandler: productHandler,
		UserHandler:    userHandler,
		PartnerHandler : partnerHandler,
		CategoryParentHandler: categoryParentHandler,
		CategoryChildHandler: *categoryChildHandler,
		
	}
}
