package main

import (
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/controller"
	"github.com/SerkanKutlu/orderService/handler"
	"github.com/SerkanKutlu/orderService/repository/mongodb"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	env := os.Getenv("GO_ENV")
	confManager := config.NewConfigurationManager("./yml", "application", env)
	//Getting Mongo Configurations
	mongoConfig := confManager.GetMongoConfiguration()
	//Getting Mongo Service
	mongoService := mongodb.GetMongoService(mongoConfig)
	//Setting Mongo Repository by using Mongo Service
	handler.SetOrderRepository(mongoService)
	//Getting Rabbit Configurations
	rabbitConfig := confManager.GetRabbitConfiguration()
	queueConfig := confManager.GetQueuesConfiguration()
	//Setting Rabbit Client
	handler.SetRabbitClient(*rabbitConfig, *queueConfig)

	//Getting Http Client Configurations
	remoteServerConfiguration := confManager.GetRemoteServerConfiguration()
	//Setting Http Client
	handler.SetHttpClient(*remoteServerConfiguration)

	//Getting Order Service that will be used at handler methods.
	orderService := handler.GetOrderService()
	orderController := controller.GetOrderController(orderService)
	productController := controller.GetProductController(orderService)
	e := echo.New()
	//Order Controls
	e.GET("/api/order", orderController.GetAllOrders)
	e.GET("/api/order/:id", orderController.GetByIdOrder)
	e.GET("/api/order/customer/:id", orderController.GetOrdersOfCustomer)
	e.POST("/api/order", orderController.PostOrder)
	e.PUT("/api/order", orderController.PutOrder)
	e.PUT("/api/order/status/:id/:status", orderController.PutOrderStatus)
	e.DELETE("/api/order/:id", orderController.DeleteOrder)
	e.DELETE("/api/order/customer/:id", orderController.DeleteOrdersOfCustomer)
	//Product Controls
	e.GET("/api/product", productController.GetAllProducts)
	e.DELETE("/api/product/:id", productController.DeleteProduct)
	e.GET("/api/product/:id", productController.GetByIdProduct)
	e.POST("/api/product", productController.PostProduct)
	e.PUT("/api/product", productController.PutProduct)
	err := e.Start(":9000")
	if err != nil {
		panic(err)
	}
}
