package config

type TopicConfig struct {
	OrderKafka OrderTopicConfig `yaml:"orderKafka"`
}

type OrderTopicConfig struct {
	OrderCreated Topic `yaml:"orderCreated"`
	OrderUpdated Topic `yaml:"orderUpdated"`
}
type Topic struct {
	Topic string `yaml:"topic"`
}
