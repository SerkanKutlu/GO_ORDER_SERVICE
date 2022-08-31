package handler

import (
	"fmt"
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/dto"
	"github.com/SerkanKutlu/orderService/events"
	"github.com/SerkanKutlu/orderService/model"
	"github.com/go-playground/validator/v10"
)

func (ds *OrderService) GetAllOrders() (*[]model.Order, *customerror.CustomError) {
	orders, err := ds.GenericRepository.GenericOrder.FindAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (ds *OrderService) GetByIdOrder(id string) (*model.Order, *customerror.CustomError) {
	order, err := ds.GenericRepository.GenericOrder.FindById(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (ds *OrderService) GetOrdersOfCustomer(customerId string) (*[]model.Order, *customerror.CustomError) {
	orders, err := ds.MongoService.FindOrdersOfCustomer(customerId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (ds *OrderService) PostOrder(orderDto *dto.OrderForCreationDto) (*string, *customerror.CustomError) {
	//Validation
	fmt.Println("post1")
	if err := validator.New().Struct(orderDto); err != nil {
		fmt.Println("post2")
		customError := customerror.NewError(err.Error(), 400)
		return nil, customError
	}
	//Dto to Order after validation passed
	newOrder := orderDto.ToOrder()
	fmt.Println("post3")
	//CustomerCheck and Address assignment
	address, err := ds.HttpClient.GetCustomerAddress(newOrder.CustomerId)
	fmt.Println("post4")
	if err != nil {
		fmt.Println("http de patladi")
		return nil, err
	}
	newOrder.Address = *address

	//Product Check, Calculate total and quantity
	var total float32
	for _, p := range newOrder.ProductIds {
		product, err := ds.GenericRepository.GenericProduct.FindById(p)
		if err != nil {
			customErr := customerror.NewError("Product with id: "+p+" not found", 404)
			return nil, customErr
		}
		total += product.Price
	}
	newOrder.Total = total
	newOrder.Quantity = len(newOrder.ProductIds)

	//Insert to db
	if err := ds.GenericRepository.GenericOrder.Insert(newOrder); err != nil {
		return nil, err
	}

	//Publish to rabbit
	fmt.Println("rabbite gidiyor")
	var event = new(events.OrderCreated)
	event.FillCreated(newOrder)
	if err := ds.RabbitClient.PublishAtCreated(event); err != nil {
		return nil, err
	}

	fmt.Println("kafkaya gidiyor")
	//Publish to kafka
	if err := ds.KafkaClient.PublishAtCreation(event); err != nil {
		fmt.Println("kafkada sorun var")
		return nil, err
	}
	//Return
	return &newOrder.Id, nil

}

func (ds *OrderService) PutOrder(orderDto *dto.OrderForUpdateDto) (*model.Order, *customerror.CustomError) {
	//Validation
	if err := validator.New().Struct(orderDto); err != nil {
		customError := customerror.NewError(err.Error(), 400)
		return nil, customError
	}

	//Product Check, Calculate total and quantity
	var total float32
	for _, p := range orderDto.ProductIds {
		product, err := ds.GenericRepository.GenericProduct.FindById(p)
		if err != nil {
			customErr := customerror.NewError("Product with id: "+product.Id+" not found", 404)
			return nil, customErr
		}
		total += product.Price
	}
	//Dto to Order after validation passed
	updatedOrder := orderDto.ToOrder()
	//Total and quantity assignment
	updatedOrder.Quantity = len(updatedOrder.ProductIds)
	updatedOrder.Total = total
	//Set createdAt field
	oldOrder, err := ds.GenericRepository.GenericOrder.FindById(updatedOrder.Id)
	if err != nil {
		return nil, err
	}
	updatedOrder.CreatedAt = oldOrder.CreatedAt

	//Set customerId field
	updatedOrder.CustomerId = oldOrder.CustomerId

	//Update order
	if err := ds.GenericRepository.GenericOrder.Update(updatedOrder, updatedOrder.Id); err != nil {
		return nil, err
	}

	//Publish to rabbit
	var event = new(events.OrderUpdated)
	event.FillUpdated(updatedOrder)
	if err := ds.RabbitClient.PublishAtUpdated(event); err != nil {
		return nil, err
	}

	fmt.Println("kafkaya gidiyor")
	//Publish to kafka
	if err := ds.KafkaClient.PublishAtUpdate(event); err != nil {
		fmt.Println("kafkada sorun var")
		return nil, err
	}
	return updatedOrder, nil

}

func (ds *OrderService) PutOrderStatus(id string, status string) *customerror.CustomError {
	if err := ds.MongoService.UpdateStatusFieldOrder(id, status); err != nil {
		return err
	}
	return nil
}

func (ds *OrderService) DeleteOrder(id string) *customerror.CustomError {
	if err := ds.GenericRepository.GenericOrder.Delete(id); err != nil {
		return err
	}
	return nil
}

func (ds *OrderService) DeleteOrdersOfCustomer(customerId string) *customerror.CustomError {

	if err := ds.MongoService.DeleteOrdersOfCustomer(customerId); err != nil {
		return err
	}
	return nil

}
