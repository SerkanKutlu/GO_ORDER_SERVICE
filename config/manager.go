package config

import (
	"github.com/spf13/viper"
)

type configurationManager struct {
	applicationConfig *ApplicationConfig
	queuesConfig      *QueueConfig
}

func NewConfigurationManager(env string) *configurationManager {
	viper.AddConfigPath("./yml")
	viper.SetConfigType("yml")
	appConfig := readApplicationConfigFile(env)
	queueConfig := readQueuesConfigFile()
	return &configurationManager{
		applicationConfig: appConfig,
		queuesConfig:      queueConfig,
	}
}

func (cm *configurationManager) GetRabbitConfiguration() *RabbitConfig {
	return &cm.applicationConfig.Rabbit
}

func (cm *configurationManager) GetMongoConfiguration() *MongoConfig {
	return &cm.applicationConfig.Mongo
}

func (cm *configurationManager) GetQueuesConfiguration() *QueueConfig {
	return cm.queuesConfig
}

func readApplicationConfigFile(env string) *ApplicationConfig {

	viper.SetConfigName("application")
	if err := viper.ReadInConfig(); err != nil {
		panic("Can not load application config file")
	}
	var appConfig ApplicationConfig
	envSub := viper.Sub(env)
	if err := envSub.Unmarshal(&appConfig); err != nil {
		panic(err.Error())
	}
	return &appConfig
}

func readQueuesConfigFile() *QueueConfig {
	viper.SetConfigName("rabbit-queue")
	if err := viper.ReadInConfig(); err != nil {
		panic("Can not load application config file")
	}
	var queueConfig QueueConfig
	envSub := viper.Sub("queue")
	if err := envSub.Unmarshal(&queueConfig); err != nil {
		panic(err.Error())
	}
	return &queueConfig
}
