package kafkaPkg

import (
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Client struct {
	ProducerConfig kafka.ConfigMap
	TopicConfig    *config.TopicConfig
	Producer       *kafka.Producer
}

func NewKafkaClient(kafkaConfig config.KafkaConfig, kafkaTopicConfig *config.TopicConfig) *Client {
	client := new(Client)
	client.TopicConfig = kafkaTopicConfig
	client.ProducerConfig = make(map[string]kafka.ConfigValue)
	for key, value := range kafkaConfig.ProducerConfig {
		client.ProducerConfig[key] = value
	}
	return client
}
func (client *Client) SetProducer() {
	producer, err := client.CreateProducer()
	if err != nil {
		panic(err.Message)
	}
	client.Producer = producer
}
func (client *Client) CreateProducer() (*kafka.Producer, *customerror.CustomError) {
	producer, err := kafka.NewProducer(&client.ProducerConfig)
	if err != nil {
		customError := customerror.NewError(err.Error(), 500)
		return nil, customError
	}
	return producer, nil
}
