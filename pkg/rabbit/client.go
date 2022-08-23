package rabbit

import (
	"edu_src_Go/config"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type client struct {
	connection  *amqp.Connection
	queueConfig *config.QueueConfig
}

func NewRabbitClient(rabbitConfig config.RabbitConfig, queueConfig config.QueueConfig) *client {
	connection := createConnection(rabbitConfig)
	return &client{
		connection:  connection,
		queueConfig: &queueConfig,
	}
}

func (client *client) SetAllConfigurations() {
	channel := client.createChannel()
	queues := client.getRegisteredQueues()
	for _, queue := range *queues {
		declareQueue(channel, queue)
		declareExchange(channel, queue)
		bindQueue(channel, queue)
	}

}

func (client *client) createChannel() *amqp.Channel {
	channel, err := client.connection.Channel()
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

func (client *client) CloseConnection() {
	err := client.connection.Close()
	if err != nil {
		panic("Rabbit mq connection close failed")
	}
}
func getConnectionUrl(rabbitConfig config.RabbitConfig) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s", rabbitConfig.Username, rabbitConfig.Password, rabbitConfig.Host, rabbitConfig.Port, rabbitConfig.VirtualHost)
}
