package kafkaPkg

import (
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Client struct {
	ProducerConfig kafka.ConfigMap
	ConsumerConfig kafka.ConfigMap
	TopicConfig    *config.TopicConfig
	Producer       *kafka.Producer
	Consumers      *[]KafkaConsumer
}

type KafkaConsumeMethod func(message *kafka.Message) error
type KafkaConsumer struct {
	Name          string
	Consumer      *kafka.Consumer
	ConsumeMethod KafkaConsumeMethod
}

func NewKafkaClient(kafkaConfig config.KafkaConfig, kafkaTopicConfig *config.TopicConfig) *Client {
	client := new(Client)
	client.TopicConfig = kafkaTopicConfig
	client.ProducerConfig = make(map[string]kafka.ConfigValue)
	client.ConsumerConfig = make(map[string]kafka.ConfigValue)
	for key, value := range kafkaConfig.ProducerConfig {
		client.ProducerConfig[key] = value
	}
	for key, value := range kafkaConfig.ConsumerConfig {
		client.ConsumerConfig[key] = value
	}
	//1297640
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
func (client *Client) CreateConsumer(topicName string) (*kafka.Consumer, *customerror.CustomError) {
	consumer, err := kafka.NewConsumer(&client.ConsumerConfig)
	if err != nil {
		customError := customerror.NewError(err.Error(), 500)
		return nil, customError
	}
	if err := consumer.Subscribe(topicName, nil); err != nil {
		customError := customerror.NewError(err.Error(), 500)
		return nil, customError
	}
	return consumer, nil
}
