package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/events"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Client struct {
	connection      *amqp.Connection
	channel         *amqp.Channel
	queueConfig     *config.QueueConfig
	publisherConfig *config.PublisherConfig
}

func (client *Client) PublishAtCreated(message *events.OrderCreated) error {
	exchangeName := client.publisherConfig.OrderPublisher.OrderCreated.ExchangeName
	routingKey := client.publisherConfig.OrderPublisher.OrderCreated.RoutingKey
	byteBody, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = client.channel.PublishWithContext(context.Background(), exchangeName, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        byteBody,
	})
	if err != nil {
		return err
	}
	return nil
}
func (client *Client) PublishAtUpdated(message *events.OrderUpdated) error {
	exchangeName := client.publisherConfig.OrderPublisher.OrderUpdated.ExchangeName
	routingKey := client.publisherConfig.OrderPublisher.OrderUpdated.RoutingKey
	byteBody, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = client.channel.PublishWithContext(context.Background(), exchangeName, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        byteBody,
	})
	if err != nil {
		return err
	}
	return nil
}
func NewRabbitClient(rabbitConfig config.RabbitConfig, queueConfig config.QueueConfig, publisherConfig config.PublisherConfig) *Client {
	connection := createConnection(rabbitConfig)
	channel := createChannel(connection)
	client := &Client{
		connection:      connection,
		channel:         channel,
		queueConfig:     &queueConfig,
		publisherConfig: &publisherConfig,
	}
	client.setAllConfigurations()
	return client
}

//Creating channel, declare queues and exchanges, binding.
func (client *Client) setAllConfigurations() {
	queues := client.getRegisteredQueues()
	for _, queue := range *queues {
		declareQueue(client.channel, queue)
		declareExchange(client.channel, queue)
		bindQueue(client.channel, queue)
	}
}

func createChannel(connection *amqp.Connection) *amqp.Channel {
	channel, err := connection.Channel()
	if err != nil {
		panic("Rabbit channel creation error: " + err.Error())
	}
	return channel
}

func declareExchange(channel *amqp.Channel, queueConfig config.Queue) {
	if err := channel.ExchangeDeclare(queueConfig.Exchange, queueConfig.ExchangeType, true, false, false, false, nil); err != nil {
		panic("Exchange declare fail: " + err.Error())
	}
}

func declareQueue(channel *amqp.Channel, queueConfig config.Queue) {
	_, err := channel.QueueDeclare(queueConfig.Queue, true, false, false, false, nil)
	if err != nil {
		panic("Queue declare fail: " + err.Error())
	}
}

func bindQueue(channel *amqp.Channel, queueConfig config.Queue) {
	if err := channel.QueueBind(queueConfig.Queue, queueConfig.RoutingKey, queueConfig.Exchange, false, nil); err != nil {
		panic("Queue binding fail: " + err.Error())
	}
}

func createConnection(rabbitConfig config.RabbitConfig) *amqp.Connection {
	amqpConfig := amqp.Config{
		Heartbeat: 30 * time.Second,
		Properties: amqp.Table{
			"connection_name": rabbitConfig.ConnectionName,
		},
	}
	connectionUrl := getConnectionUrl(rabbitConfig)
	connection, err := amqp.DialConfig(connectionUrl, amqpConfig)
	if err != nil {
		_ = connection.Close()
		panic("Rabbit mq connection failed")
	}
	fmt.Println("Rabbit connection is done")
	return connection
}

func (client *Client) CloseConnection() {
	err := client.connection.Close()
	if err != nil {
		panic("Rabbit mq connection close failed")
	}
}

func getConnectionUrl(rabbitConfig config.RabbitConfig) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s", rabbitConfig.Username, rabbitConfig.Password, rabbitConfig.Host, rabbitConfig.Port, rabbitConfig.VirtualHost)
}
