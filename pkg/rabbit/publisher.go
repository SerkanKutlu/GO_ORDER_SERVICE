package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (client *Client) PublishAtCreated(message *events.OrderCreated) *customerror.CustomError {
	exchangeName := client.QueueConfig.Order.OrderCreated.Exchange
	routingKey := client.QueueConfig.Order.OrderCreated.RoutingKey
	byteBody, err := json.Marshal(message)
	if err != nil {
		return customerror.InternalServerError
	}
	err = client.Channel.PublishWithContext(context.Background(), exchangeName, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        byteBody,
	})
	if err != nil {
		return customerror.InternalServerError
	}
	return nil
}
func (client *Client) PublishAtUpdated(message *events.OrderUpdated) *customerror.CustomError {
	exchangeName := client.QueueConfig.Order.OrderUpdated.Exchange
	routingKey := client.QueueConfig.Order.OrderUpdated.RoutingKey
	byteBody, err := json.Marshal(message)
	if err != nil {
		return customerror.InternalServerError
	}
	err = client.Channel.PublishWithContext(context.Background(), exchangeName, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        byteBody,
	})
	if err != nil {
		return customerror.InternalServerError
	}
	return nil
}
func (client *Client) EnsureConnection() {
	go func() {
		<-client.ErrorChannel
		fmt.Println("Retrying to connect rabbit mq")
		connectionError := customerror.InternalServerError
		for connectionError != nil {
			connectionError = client.ReConnect()
		}
		client.EnsureConnection()
	}()
}
