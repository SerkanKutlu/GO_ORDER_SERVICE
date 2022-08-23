package handler

import (
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/dataservice"
	"github.com/SerkanKutlu/orderService/pkg/rabbit"
)

type DataAccessService struct {
	OrderService   dataservice.OrderDataInterface
	ProductService dataservice.ProductDataInterface
	RabbitClient   *rabbit.Client
}

var dataAccessService = new(DataAccessService)
var rabbitClient *rabbit.Client

func GetDataServices(env string) *DataAccessService {
	if rabbitClient != nil {
		dataAccessService.RabbitClient = rabbitClient
		return dataAccessService
	}
	//RABBIT MQ
	confManager := config.NewConfigurationManager("./yml", "application", env)
	rabbitConfig := confManager.GetRabbitConfiguration()
	queueConfig := confManager.GetQueuesConfiguration()
	publisherConfig := confManager.GetPublisherConfiguration()
	rabbitClient := rabbit.NewRabbitClient(*rabbitConfig, *queueConfig, *publisherConfig)
	dataAccessService.RabbitClient = rabbitClient
	return dataAccessService
}

func SetOrderDataService(orderService dataservice.OrderDataInterface) {
	dataAccessService.OrderService = orderService
}

func SetProductDataService(productService dataservice.ProductDataInterface) {
	dataAccessService.ProductService = productService
}
