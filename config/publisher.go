package config

type PublisherConfig struct {
	Order OrderPublisher `yaml:"order"`
}
type OrderPublisher struct {
	OrderCreated Publisher `yaml:"orderCreated"`
	OrderUpdated Publisher `yaml:"orderUpdated"`
}
type Publisher struct {
	Exchange   string `yaml:"exchange"`
	RoutingKey string `yaml:"routingKey"`
}
