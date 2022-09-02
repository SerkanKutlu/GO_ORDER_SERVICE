package config

type TopicConfig struct {
	OrderKafka OrderTopicConfig `yaml:"orderKafka"`
}

type OrderTopicConfig struct {
	OrderCreated    Topic `yaml:"orderCreated"`
	OrderUpdated    Topic `yaml:"orderUpdated"`
	OrderCreated4sn Topic `yaml:"orderCreated4Sn"`
	OrderCreated8sn Topic `yaml:"orderCreated8Sn"`
	OrderCreatedDLQ Topic `yaml:"orderCreatedDLQ"`
}
type Topic struct {
	Topic string `yaml:"topic"`
}
