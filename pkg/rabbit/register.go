package rabbit

import (
	"github.com/SerkanKutlu/orderService/config"
)

var registeredQueues []config.Queue

func (client *Client) getRegisteredQueues() *[]config.Queue {
	if registeredQueues != nil {
		return &registeredQueues
	}
	registeredQueues = append(registeredQueues, client.queueConfig.Order.OrderCreated)
	registeredQueues = append(registeredQueues, client.queueConfig.Order.OrderUpdated)
	return &registeredQueues
}
