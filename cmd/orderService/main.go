package main

import (
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/controller"
	"github.com/SerkanKutlu/orderService/handler"
	"github.com/SerkanKutlu/orderService/middleware"
	"github.com/SerkanKutlu/orderService/repository/mongodb"
	"github.com/labstack/echo/v4"
	_ "net/http/pprof"
)

func main() {
	env := "dev"
	confManager := config.NewConfigurationManager("yml", "application", env)
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

	//Getting Kafka Configurations
	kafkaConfig := confManager.GetKafkaConfiguration()
	topicConfig := confManager.GetKafkaTopicConfiguration()
	//Setting Kafka Client
	kafkaClient := handler.SetKafkaClient(*kafkaConfig, *topicConfig)
	kafkaClient.SetProducer()
	//kafkaClient.PublishBigMessage()
	//Setting redis client
	handler.SetRedisClient(confManager.GetRedisConfiguration())

	//Getting Order Service that will be used at handler methods.
	orderService := handler.GetOrderService()
	orderController := controller.GetOrderController(orderService)
	productController := controller.GetProductController(orderService)
	e := echo.New()
	e.Use(middleware.ErrorHandler)
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
	e.GET("/api/order/large", orderController.LargePush)
	err := e.Start(":4001")
	if err != nil {
		panic(err)
	}
}
