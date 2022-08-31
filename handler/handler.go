package handler

import (
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/model"
	"github.com/SerkanKutlu/orderService/pkg/kafka"
	"github.com/SerkanKutlu/orderService/pkg/rabbit"
	"github.com/SerkanKutlu/orderService/pkg/utils"
	"github.com/SerkanKutlu/orderService/repository/mongodb"
)

type OrderService struct {
	MongoService      *mongodb.MongoService
	RabbitClient      *rabbit.Client
	GenericRepository *mongodb.GenericRepository
	HttpClient        *utils.HttpClient
	KafkaClient       *kafka.Client
}

var orderService = new(OrderService)

func GetOrderService() *OrderService {
	return orderService
}
func SetOrderRepository(mongoService *mongodb.MongoService) {
	orderService.MongoService = mongoService
	orderService.GenericRepository = new(mongodb.GenericRepository)
	orderService.GenericRepository.GenericOrder = &mongodb.GenericCollection[model.Order]{
		Collection: mongoService.Orders,
	}
	orderService.GenericRepository.GenericProduct = &mongodb.GenericCollection[model.Product]{
		Collection: mongoService.Products,
	}
}
func SetRabbitClient(rabbitConfig config.RabbitConfig, queueConfig config.QueueConfig) {
	orderService.RabbitClient = new(rabbit.Client)
	orderService.RabbitClient = rabbit.NewRabbitClient(rabbitConfig, queueConfig)
	orderService.RabbitClient.EnsureConnection()
}

func SetHttpClient(remoteServers config.RemoteServicesConfig) {
	orderService.HttpClient = new(utils.HttpClient)
	orderService.HttpClient.ServiceUrlMap = make(map[string]string)
	orderService.HttpClient.ServiceUrlMap[remoteServers.CustomerService.Name] = remoteServers.CustomerService.BaseUrl
}

func SetKafkaClient(kafkaConfig config.KafkaConfig, topicConfig config.TopicConfig) {
	orderService.KafkaClient = new(kafka.Client)
	orderService.KafkaClient = kafka.NewKafkaClient(kafkaConfig, &topicConfig)
}
