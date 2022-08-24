package main

import (
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/handler"
	"github.com/SerkanKutlu/orderService/repository/mongodb"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	env := os.Getenv("GO_ENV")

	confManager := config.NewConfigurationManager("./yml", "application", env)

	//DATA ACCESS
	mongoConfig := confManager.GetMongoConfiguration()
	mongoService := mongodb.GetMongoService(mongoConfig)
	handler.SetDataServices(mongoService)
	//handler.SetProductDataService(mongoService)
	//handler.SetOrderDataService(mongoService)
	orderService := handler.GetDataServices(env)

	e := echo.New()
	//Order Controls
	e.GET("/api/order", orderService.GetAllOrders)
	e.GET("/api/order/:id", orderService.GetByIdOrder)
	e.GET("/api/order/customer/:id", orderService.GetOrdersOfCustomers)
	e.POST("/api/order", orderService.PostOrder)
	e.PUT("/api/order", orderService.PutOrder)
	e.PUT("/api/order/status/:id/:status", orderService.PutOrderStatus)
	e.DELETE("/api/order/:id", orderService.DeleteOrder)
	e.DELETE("/api/order/customer/:id", orderService.DeleteOrdersOfCustomer)
	//Product Controls
	e.GET("/api/product", orderService.GetAllProducts)
	e.DELETE("/api/product/:id", orderService.DeleteProduct)
	e.GET("/api/product/:id", orderService.GetByIdProduct)
	e.POST("/api/product", orderService.PostProduct)
	e.PUT("/api/product", orderService.PutProduct)
	err := e.Start(":9000")
	if err != nil {
		panic(err)
	}
}
