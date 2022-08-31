package kafka

import (
	"encoding/json"
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/events"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (client *Client) PublishAtCreation(message *events.OrderCreated) *customerror.CustomError {
	deliveryChan := make(chan kafka.Event, 10000)
	topicPartition := &kafka.TopicPartition{
		Topic:     &client.TopicConfig.OrderKafka.OrderCreated.Topic,
		Partition: kafka.PartitionAny,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return customerror.NewError(err, 500)
	}
	kafkaMessage := &kafka.Message{
		TopicPartition: *topicPartition,
		Value:          messageBytes,
	}
	err = client.Producer.Produce(kafkaMessage, deliveryChan)
	if err != nil {
		return customerror.NewError(err, 500)
	}
	return nil
}
func (client *Client) PublishAtUpdate(message *events.OrderUpdated) *customerror.CustomError {
	deliveryChan := make(chan kafka.Event)
	topicPartition := &kafka.TopicPartition{
		Topic:     &client.TopicConfig.OrderKafka.OrderUpdated.Topic,
		Partition: kafka.PartitionAny,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return customerror.NewError(err, 500)
	}
	kafkaMessage := &kafka.Message{
		TopicPartition: *topicPartition,
		Value:          messageBytes,
	}
	err = client.Producer.Produce(kafkaMessage, deliveryChan)
	if err != nil {
		return customerror.NewError(err, 500)
	}
	return nil
}
