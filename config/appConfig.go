package config

type ApplicationConfig struct {
	Rabbit RabbitConfig `yaml:"rabbit"`
	Mongo  MongoConfig  `yaml:"mongo"`
	Kafka  KafkaConfig  `yaml:"kafka"`
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

type KafkaConfig struct {
	ProducerConfig map[string]string `yaml:"producerConfig"`
	ConsumerConfig map[string]string `yaml:"consumerConfig"`
}
