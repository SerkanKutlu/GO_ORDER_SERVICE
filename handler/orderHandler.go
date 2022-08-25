package handler

import (
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/dto"
	"github.com/SerkanKutlu/orderService/events"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (ds *OrderService) GetAllOrders(c echo.Context) error {
	orders, err := ds.GenericRepository.GenericOrder.FindAll()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, orders)
}

func (ds *OrderService) GetByIdOrder(c echo.Context) error {
	id := c.Param("id")
	order, err := ds.GenericRepository.GenericOrder.FindById(id)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, order)
}

func (ds *OrderService) GetOrdersOfCustomers(c echo.Context) error {
	customerId := c.Param("id")
	orders, err := ds.MongoService.FindOrdersOfCustomer(customerId)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, orders)
}

func (ds *OrderService) PostOrder(c echo.Context) error {
	var orderDto dto.OrderForCreationDto
	//Binding
	if err := c.Bind(&orderDto); err != nil {
		return c.JSON(customerror.InvalidBodyError.StatusCode, customerror.InvalidBodyError)
	}

	//Validation
	if err := validator.New().Struct(orderDto); err != nil {
		customError := customerror.NewError(err.Error(), 400)
		return c.JSON(customError.StatusCode, customError)
	}
	//Dto to Order after validation passed
	newOrder := orderDto.ToOrder()

	//CustomerCheck and Address assignment
	address, err := ds.HttpClient.GetCustomerAddress(newOrder.CustomerId)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	newOrder.Address = *address

	//Product Check, Calculate total and quantity
	var total float32
	for _, p := range newOrder.ProductIds {
		product, err := ds.GenericRepository.GenericProduct.FindById(p)
		if err != nil {
			customErr := customerror.NewError("Product with id: "+p+" not found", 404)
			return c.JSON(customErr.StatusCode, customErr)
		}
		total += product.Price
	}
	newOrder.Total = total
	newOrder.Quantity = len(newOrder.ProductIds)

	//Insert to db
	if err := ds.GenericRepository.GenericOrder.Insert(newOrder); err != nil {
		return c.JSON(err.StatusCode, err)
	}

	//Publish to rabbit
	var event = new(events.OrderCreated)
	event.FillCreated(newOrder)
	if err := ds.RabbitClient.PublishAtCreated(event); err != nil {
		return c.JSON(err.StatusCode, err)
	}

	//Return
	return c.JSON(201, newOrder.Id)

}

func (ds *OrderService) PutOrder(c echo.Context) error {
	var orderDto dto.OrderForUpdateDto
	//Binding
	if err := c.Bind(&orderDto); err != nil {
		return c.JSON(customerror.InvalidBodyError.StatusCode, customerror.InvalidBodyError)
	}

	//Validation
	if err := validator.New().Struct(orderDto); err != nil {
		customError := customerror.NewError(err.Error(), 400)
		return c.JSON(customError.StatusCode, customError)
	}

	//Product Check, Calculate total and quantity
	var total float32
	for _, p := range orderDto.ProductIds {
		product, err := ds.GenericRepository.GenericProduct.FindById(p)
		if err != nil {
			customErr := customerror.NewError("Product with id: "+product.Id+" not found", 404)
			return c.JSON(customErr.StatusCode, customErr)
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
		return c.JSON(err.StatusCode, err)
	}
	updatedOrder.CreatedAt = oldOrder.CreatedAt

	//Set customerId field
	updatedOrder.CustomerId = oldOrder.CustomerId

	//Update order
	if err := ds.GenericRepository.GenericOrder.Update(updatedOrder, updatedOrder.Id); err != nil {
		return c.JSON(err.StatusCode, err)
	}

	//Publish to rabbit
	var event = new(events.OrderUpdated)
	event.FillUpdated(updatedOrder)
	if err := ds.RabbitClient.PublishAtUpdated(event); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")

}

func (ds *OrderService) PutOrderStatus(c echo.Context) error {
	id := c.Param("id")
	status := c.Param("status")
	if status == "" {
		err := customerror.NewError("Status can not be empty", 400)
		return c.JSON(err.StatusCode, err)
	}
	if err := ds.MongoService.UpdateStatusFieldOrder(id, status); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}

func (ds *OrderService) DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	if err := ds.GenericRepository.GenericOrder.Delete(id); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}

func (ds *OrderService) DeleteOrdersOfCustomer(c echo.Context) error {
	customerId := c.Param("id")
	if err := ds.MongoService.DeleteOrdersOfCustomer(customerId); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")

}
