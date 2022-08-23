package handler

import (
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/dto"
	"github.com/SerkanKutlu/orderService/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (ds *DataAccessService) GetAllOrders(c echo.Context) error {
	orders, err := ds.OrderService.FindAllOrders()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, orders)
}

func (ds *DataAccessService) GetByIdOrder(c echo.Context) error {
	id := c.Param("id")
	order, err := ds.OrderService.FindByIdOrder(id)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, order)
}

func (ds *DataAccessService) GetOrdersOfCustomers(c echo.Context) error {
	customerId := c.Param("id")
	orders, err := ds.OrderService.FindOrdersOfCustomer(customerId)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, orders)
}

func (ds *DataAccessService) PostOrder(c echo.Context) error {
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
	address, err := utils.GetCustomerAddress(newOrder.CustomerId)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	newOrder.Address = *address

	//Product Check, Calculate total and quantity
	var total float32
	for _, p := range newOrder.ProductIds {
		product, err := ds.ProductService.FindByIdProduct(p)
		if err != nil {
			customErr := customerror.NewError("Product with id: "+p+" not found", 404)
			return c.JSON(customErr.StatusCode, customErr)
		}
		total += product.Price
	}
	newOrder.Total = total
	newOrder.Quantity = len(newOrder.ProductIds)

	//Insert to db
	if err := ds.OrderService.InsertOrder(newOrder); err != nil {
		return c.JSON(err.StatusCode, err)
	}

	//Publish to rabbit
	//event := new(events.OrderCreated)
	//event.FillCreated(newOrder)
	//rabbit.Publish(event, "go.create")

	//Return
	return c.JSON(201, newOrder.Id)

}

func (ds *DataAccessService) PutOrder(c echo.Context) error {
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
		product, err := ds.ProductService.FindByIdProduct(p)
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
	oldOrder, err := ds.OrderService.FindByIdOrder(updatedOrder.Id)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	updatedOrder.CreatedAt = oldOrder.CreatedAt

	//Set customerId field
	updatedOrder.CustomerId = oldOrder.CustomerId

	//Update order
	if err := ds.OrderService.UpdateOrder(updatedOrder); err != nil {
		return c.JSON(err.StatusCode, err)
	}

	//Publish to rabbit
	//event := new(events.OrderUpdated)
	//event.FillUpdated(updatedOrder)
	//rabbit.Publish(event, "go.update")
	//Return
	return c.JSON(200, "")

}

func (ds *DataAccessService) PutOrderStatus(c echo.Context) error {
	id := c.Param("id")
	status := c.Param("status")
	if status == "" {
		err := customerror.NewError("Status can not be empty", 400)
		return c.JSON(err.StatusCode, err)
	}
	if err := ds.OrderService.UpdateStatusFieldOrder(id, status); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}

func (ds *DataAccessService) DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	if err := ds.OrderService.DeleteOrder(id); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}

func (ds *DataAccessService) DeleteOrdersOfCustomer(c echo.Context) error {
	customerId := c.Param("id")
	if err := ds.OrderService.DeleteOrdersOfCustomer(customerId); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")

}
