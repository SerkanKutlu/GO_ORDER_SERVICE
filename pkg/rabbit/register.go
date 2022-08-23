package rabbit

import (
	"github.com/SerkanKutlu/orderService/config"
)

var registeredQueues []config.Queue

func (client *Client) GetRegisteredQueues() *[]config.Queue {
	if registeredQueues != nil {
		return &registeredQueues
	}
	registeredQueues = append(registeredQueues, client.QueueConfig.Order.OrderCreated)
	registeredQueues = append(registeredQueues, client.QueueConfig.Order.OrderUpdated)
	return &registeredQueues
}
