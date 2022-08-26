package controller

import (
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/dto"
	"github.com/SerkanKutlu/orderService/handler"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	Controller *Controller
}

func GetProductController(orderService *handler.OrderService) *ProductController {
	return &ProductController{Controller: &Controller{OrderService: orderService}}

}

func (controller *ProductController) GetAllProducts(c echo.Context) error {
	products, err := controller.Controller.OrderService.GetAllProducts()
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, products)
}
func (controller *ProductController) GetByIdProduct(c echo.Context) error {
	id := c.Param("id")
	product, err := controller.Controller.OrderService.GetByIdProduct(id)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, product)
}
func (controller *ProductController) PostProduct(c echo.Context) error {
	var productDto *dto.ProductForCreationDto
	//Binding
	if err := c.Bind(&productDto); err != nil {
		return c.JSON(customerror.InvalidBodyError.StatusCode, customerror.InvalidBodyError)
	}
	createdId, err := controller.Controller.OrderService.PostProduct(productDto)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(201, createdId)
}
func (controller *ProductController) PutProduct(c echo.Context) error {
	var productDto *dto.ProductForUpdateDto
	//Binding
	if err := c.Bind(&productDto); err != nil {
		return c.JSON(customerror.InvalidBodyError.StatusCode, customerror.InvalidBodyError)
	}
	updatedProduct, err := controller.Controller.OrderService.PutProduct(productDto)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, updatedProduct)
}
func (controller *ProductController) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	err := controller.Controller.OrderService.DeleteProduct(id)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}
