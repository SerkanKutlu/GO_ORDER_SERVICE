package config

type QueueConfig struct {
	Order OrderQueueConfig `yaml:"order"`
}
type OrderQueueConfig struct {
	OrderCreated Queue `yaml:"orderCreated"`
	OrderUpdated Queue `yaml:"orderUpdated"`
}
type Queue struct {
	Exchange     string `yaml:"exchange"`
	ExchangeType string `yaml:"exchangeType"`
	RoutingKey   string `yaml:"routingKey"`
	Queue        string `yaml:"queue"`
}
