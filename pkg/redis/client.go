package redisPkg

import (
	"context"
	"encoding/json"
	"github.com/SerkanKutlu/orderService/config"
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/events"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	redisConfig *config.RedisConfig
	client      *redis.Client
}

func NewRedisClient(redisConfig *config.RedisConfig) *Client {

	client := new(Client)
	client.redisConfig = redisConfig
	redisOpts := &redis.Options{Addr: redisConfig.ConnectionString}
	client.client = redis.NewClient(redisOpts)
	return client
}

func (c *Client) PublishMain(orderCreated *events.OrderCreated) *customerror.CustomError {

	payload, _ := json.Marshal(orderCreated)
	if err := c.client.Publish(context.Background(), c.redisConfig.Channel, payload).Err(); err != nil {
		customerror.NewError("problem at redis", 500)
	}
	return nil
}
