package rabbit

import (
	"errors"
	"fmt"
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/customerror"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	Connection   *amqp.Connection
	Channel      *amqp.Channel
	QueueConfig  *config.QueueConfig
	RabbitConfig *config.RabbitConfig
	ErrorChannel chan error
}

func NewRabbitClient(rabbitConfig config.RabbitConfig, queueConfig config.QueueConfig) *Client {
	client := &Client{
		QueueConfig:  &queueConfig,
		RabbitConfig: &rabbitConfig,
	}
	connection, err := client.createConnection(rabbitConfig)
	if err != nil {
		panic("Rabbit mq client could not be created")
	}
	channel := client.createChannel(connection)
	client.Connection = connection
	client.Channel = channel
	client.ErrorChannel = make(chan error)
	client.setAllConfigurations()
	return client
}

func (client *Client) ReConnect() *customerror.CustomError {
	newConnection, err := client.createConnection(*client.RabbitConfig)
	if err != nil {
		return err
	}
	newChannel := client.createChannel(newConnection)
	client.Connection = newConnection
	client.Channel = newChannel
	return nil
}

// Creating channel, declare queues and exchanges, binding.
func (client *Client) setAllConfigurations() {
	queues := client.GetRegisteredQueues()
	for _, queue := range *queues {
		declareQueue(client.Channel, queue)
		declareExchange(client.Channel, queue)
		bindQueue(client.Channel, queue)
	}
}

func (client *Client) createChannel(connection *amqp.Connection) *amqp.Channel {
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

func (client *Client) createConnection(rabbitConfig config.RabbitConfig) (*amqp.Connection, *customerror.CustomError) {
	amqpConfig := amqp.Config{
		Properties: amqp.Table{
			"connection_name": rabbitConfig.ConnectionName,
		},
	}
	connectionUrl := getConnectionUrl(rabbitConfig)
	connection, err := amqp.DialConfig(connectionUrl, amqpConfig)
	if err != nil {
		return nil, customerror.InternalServerError
	}
	go func() {
		<-connection.NotifyClose(make(chan *amqp.Error))
		client.ErrorChannel <- errors.New("rabbit connection is down")
	}()

	return connection, nil
}

func (client *Client) CloseConnection() {
	err := client.Connection.Close()
	if err != nil {
		panic("Rabbit mq connection close failed")
	}
}

func getConnectionUrl(rabbitConfig config.RabbitConfig) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s", rabbitConfig.Username, rabbitConfig.Password, rabbitConfig.Host, rabbitConfig.Port, rabbitConfig.VirtualHost)
}
