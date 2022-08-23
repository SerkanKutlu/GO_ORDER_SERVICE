package config

type ApplicationConfig struct {
	Rabbit RabbitConfig `yaml:"rabbit"`
	Mongo  MongoConfig  `yaml:"mongo"`
}
type RabbitConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	VirtualHost    string `yaml:"virtualHost"`
	ConnectionName string `yaml:"connectionName"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
}

type MongoConfig struct {
	ConnectionString string            `yaml:"connectionString"`
	Database         string            `yaml:"database"`
	Collection       map[string]string `yaml:"collection"`
}

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
