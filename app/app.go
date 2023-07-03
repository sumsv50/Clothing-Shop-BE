package app

import (
	"clothing-shop/app/config"
	"clothing-shop/handler"
	"clothing-shop/service"
)

type ApplicationContext struct {
	ProductHandler *handler.ProductHandler
}

func NewApp() *ApplicationContext {
	db := config.ConnectToDB()
	productService := service.NewProductService(db)
	productHandler := handler.NewProductHandler(*productService)

	return &ApplicationContext{ProductHandler: productHandler}
}
