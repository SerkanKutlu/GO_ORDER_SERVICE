package kafka

import (
	"github.com/SerkanKutlu/orderService/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Client struct {
	KafkaConfig kafka.ConfigMap
}

func NewKafkaClient(kafkaConfig config.KafkaConfig, kafkaTopicConfig config.TopicConfig) *Client {
	client := new(Client)
	for key, value := range kafkaConfig.ConfigMap {
		err := client.KafkaConfig.SetKey(key, value)
		if err != nil {
			panic("Fix the error of mapping kafka configurations")
		}
	}
	return client

}
