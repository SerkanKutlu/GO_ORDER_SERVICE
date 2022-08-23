package main

import (
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/dataservice/mongodb"
	"github.com/SerkanKutlu/orderService/handler"
	"github.com/SerkanKutlu/orderService/pkg/rabbit"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {

	env := os.Getenv("GO_ENV")

	confManager := config.NewConfigurationManager(env)

	//RABBIT MQ
	rabbitConfig := confManager.GetRabbitConfiguration()
	queueConfig := confManager.GetQueuesConfiguration()
	rabbitClient := rabbit.NewRabbitClient(*rabbitConfig, *queueConfig)
	rabbitClient.SetAllConfigurations()
	defer rabbitClient.CloseConnection()
	///MONGO
	mongoConfig := confManager.GetMongoConfiguration()
	mongo := mongodb.GetMongoDb(mongoConfig)
	handler.SetProductDataService(mongo)
	handler.SetOrderDataService(mongo)
	orderServiceHandler := handler.GetDataServices()

	e := echo.New()
	//Order Controls
	e.GET("/api/order", orderServiceHandler.GetAllOrders)
	e.GET("/api/order/:id", orderServiceHandler.GetByIdOrder)
	e.GET("/api/order/customer/:id", orderServiceHandler.GetOrdersOfCustomers)
	e.POST("/api/order", orderServiceHandler.PostOrder)
	e.PUT("/api/order", orderServiceHandler.PutOrder)
	e.PUT("/api/order/status/:id/:status", orderServiceHandler.PutOrderStatus)
	e.DELETE("/api/order/:id", orderServiceHandler.DeleteOrder)
	e.DELETE("/api/order/customer/:id", orderServiceHandler.DeleteOrdersOfCustomer)
	//Product Controls
	e.GET("/api/product", orderServiceHandler.GetAllProducts)
	e.DELETE("/api/product/:id", orderServiceHandler.DeleteProduct)
	e.GET("/api/product/:id", orderServiceHandler.GetByIdProduct)
	e.POST("/api/product", orderServiceHandler.PostProduct)
	e.PUT("/api/product", orderServiceHandler.PutProduct)

	err := e.Start(":9000")
	if err != nil {
		panic(err)
	}
}
