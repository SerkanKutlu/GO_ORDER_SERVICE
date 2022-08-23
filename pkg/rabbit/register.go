package rabbit

import "edu_src_Go/config"

var registeredQueues []config.Queue

func (client *client) getRegisteredQueues() *[]config.Queue {
	if registeredQueues != nil {
		return &registeredQueues
	}
	registeredQueues = append(registeredQueues, client.queueConfig.Order.OrderCreated)
	registeredQueues = append(registeredQueues, client.queueConfig.Order.OrderUpdated)
	return &registeredQueues
}
