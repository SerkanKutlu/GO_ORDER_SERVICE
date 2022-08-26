package controller

import (
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/dto"
	"github.com/SerkanKutlu/orderService/handler"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Controller *Controller
}

func GetOrderController(orderService *handler.OrderService) *OrderController {
	return &OrderController{Controller: &Controller{OrderService: orderService}}

}

func (controller *OrderController) GetAllOrders(c echo.Context) error {
	orders, err := controller.Controller.OrderService.GetAllOrders()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, orders)
}
func (controller *OrderController) GetByIdOrder(c echo.Context) error {
	id := c.Param("id")
	order, err := controller.Controller.OrderService.GetByIdOrder(id)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, order)
}
func (controller *OrderController) GetOrdersOfCustomer(c echo.Context) error {
	customerId := c.Param("id")
	orders, err := controller.Controller.OrderService.GetOrdersOfCustomer(customerId)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, orders)
}
func (controller *OrderController) PostOrder(c echo.Context) error {
	var orderDto dto.OrderForCreationDto
	//Binding
	if err := c.Bind(&orderDto); err != nil {
		return c.JSON(customerror.InvalidBodyError.StatusCode, customerror.InvalidBodyError)
	}
	createdId, err := controller.Controller.OrderService.PostOrder(&orderDto)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(201, createdId)
}
func (controller *OrderController) PutOrder(c echo.Context) error {
	var orderDto *dto.OrderForUpdateDto
	//Binding
	if err := c.Bind(&orderDto); err != nil {
		return c.JSON(customerror.InvalidBodyError.StatusCode, customerror.InvalidBodyError)
	}
	updatedOrder, err := controller.Controller.OrderService.PutOrder(orderDto)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, updatedOrder)
}
func (controller *OrderController) PutOrderStatus(c echo.Context) error {
	id := c.Param("id")
	status := c.Param("status")
	if status == "" {
		err := customerror.NewError("Status can not be empty", 400)
		return c.JSON(err.StatusCode, err)
	}
	if err := controller.Controller.OrderService.PutOrderStatus(id, status); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}
func (controller *OrderController) DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	if err := controller.Controller.OrderService.DeleteOrder(id); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}
func (controller *OrderController) DeleteOrdersOfCustomer(c echo.Context) error {
	customerId := c.Param("id")
	err := controller.Controller.OrderService.DeleteOrdersOfCustomer(customerId)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}
