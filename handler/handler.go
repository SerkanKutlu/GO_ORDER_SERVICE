package handler

import (
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/model"
	"github.com/SerkanKutlu/orderService/pkg/rabbit"
	"github.com/SerkanKutlu/orderService/repository/mongodb"
)

type OrderService struct {
	MongoService *mongodb.MongoService
	//OrderService   repository.IOrderRepository
	//ProductService repository.IProductRepository
	RabbitClient *rabbit.Client
}
type GenericRepository struct {
	GenericOrder   *mongodb.GenericRepository[model.Order]
	GenericProduct *mongodb.GenericRepository[model.Product]
}

var genericRepository *GenericRepository
var dataAccessService = new(OrderService)
var rabbitClient *rabbit.Client

func GetDataServices(env string) *OrderService {
	if rabbitClient != nil {
		dataAccessService.RabbitClient = rabbitClient
		return dataAccessService
	}
	//RABBIT MQ
	confManager := config.NewConfigurationManager("./yml", "application", env)
	rabbitConfig := confManager.GetRabbitConfiguration()
	queueConfig := confManager.GetQueuesConfiguration()
	rabbitClient := rabbit.NewRabbitClient(*rabbitConfig, *queueConfig)
	dataAccessService.RabbitClient = rabbitClient
	return dataAccessService
}
func SetDataServices(mongoService *mongodb.MongoService) {
	dataAccessService.MongoService = mongoService
	genericRepository = new(GenericRepository)
	genericRepository.GenericOrder = &mongodb.GenericRepository[model.Order]{
		Collection: mongoService.Orders,
	}
	genericRepository.GenericProduct = &mongodb.GenericRepository[model.Product]{
		Collection: mongoService.Products,
	}
}

//func SetOrderDataService(orderService repository.IOrderRepository) {
//	dataAccessService.OrderService = orderService
//}
//
//func SetProductDataService(productService repository.IProductRepository) {
//	dataAccessService.ProductService = productService
//}
