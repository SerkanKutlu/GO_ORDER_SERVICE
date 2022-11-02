package kafkaPkg

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
func (client *Client) PublishLargeFile(message any) *customerror.CustomError {
	deliveryChan := make(chan kafka.Event)
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
func (client *Client) Publish(message *[]byte, partitionNumber int32, pkgName string, isEnd bool) {
	deliveryChan := make(chan kafka.Event, 10000)
	topic := "bigDataHere"
	topicPartition := &kafka.TopicPartition{
		Topic:     &topic,
		Partition: partitionNumber,
	}
	kafkaMessage := &kafka.Message{
		TopicPartition: *topicPartition,
		Value:          *message,
	}

	if isEnd == true {
		AddHeaderToMessage(kafkaMessage, "pkgName", pkgName+"end")
	} else {
		AddHeaderToMessage(kafkaMessage, "pkgName", pkgName)
	}
	err := client.Producer.Produce(kafkaMessage, deliveryChan)
	if err != nil {
		panic(err)
	}
}

func AddHeaderToMessage(message *kafka.Message, headerKey string, headerValue string) {
	newHeader := kafka.Header{
		Key:   headerKey,
		Value: []byte(headerValue),
	}
	//If already exist header
	for index, header := range message.Headers {
		if header.Key == headerKey {
			message.Headers[index] = newHeader
			return
		}
	}
	message.Headers = append(message.Headers, newHeader)
}
