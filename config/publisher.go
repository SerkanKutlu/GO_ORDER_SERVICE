package config

type PublisherConfig struct {
	OrderPublisher OrderPublisher `yaml:"orderPublisher"`
}
type OrderPublisher struct {
	OrderCreated Publisher `yaml:"orderCreated"`
	OrderUpdated Publisher `yaml:"orderUpdated"`
}
type Publisher struct {
	ExchangeName string `yaml:"exchangeName"`
	RoutingKey   string `yaml:"routingKey"`
}
