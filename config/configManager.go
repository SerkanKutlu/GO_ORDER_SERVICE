package config

import (
	"github.com/spf13/viper"
)

type configurationManager struct {
	applicationConfig    *ApplicationConfig
	queuesConfig         *QueueConfig
	remoteServicesConfig *RemoteServicesConfig
}

func NewConfigurationManager(path string, file string, env string) *configurationManager {
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	appConfig := readApplicationConfigFile(env, file)
	queueConfig := readQueuesConfigFile(env, file)
	remoteServerConfig := readRemoteServicesConfigFile(env, file)
	return &configurationManager{
		applicationConfig:    appConfig,
		queuesConfig:         queueConfig,
		remoteServicesConfig: remoteServerConfig,
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

func (cm *configurationManager) GetRemoteServerConfiguration() *RemoteServicesConfig {
	return cm.remoteServicesConfig
}

func readApplicationConfigFile(env string, file string) *ApplicationConfig {

	viper.SetConfigName(file)
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
func readQueuesConfigFile(env string, file string) *QueueConfig {
	viper.SetConfigName(file)
	if err := viper.ReadInConfig(); err != nil {
		panic("Can not load application config file")
	}
	var queueConfig QueueConfig
	envSub := viper.Sub(env)
	if err := envSub.Unmarshal(&queueConfig); err != nil {
		panic(err.Error())
	}
	return &queueConfig
}
func readRemoteServicesConfigFile(env string, file string) *RemoteServicesConfig {
	viper.SetConfigName(file)
	if err := viper.ReadInConfig(); err != nil {
		panic("Can not load application config file")
	}
	var remoteConfig RemoteServicesConfig
	envSub := viper.Sub(env)
	if err := envSub.Unmarshal(&remoteConfig); err != nil {
		panic(err.Error())
	}
	return &remoteConfig
}
