package handler

import (
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (ds *OrderService) GetAllProducts(c echo.Context) error {
	products, err := ds.GenericRepository.GenericProduct.FindAll()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, products)
}

func (ds *OrderService) GetByIdProduct(c echo.Context) error {
	id := c.Param("id")
	product, err := ds.GenericRepository.GenericProduct.FindById(id)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, product)
}

func (ds *OrderService) PostProduct(c echo.Context) error {
	var productDto dto.ProductForCreationDto
	//Binding
	if err := c.Bind(&productDto); err != nil {
		return c.JSON(customerror.InvalidBodyError.StatusCode, customerror.InvalidBodyError)
	}

	//Validation
	if err := validator.New().Struct(productDto); err != nil {
		customError := customerror.NewError(err.Error(), 400)
		return c.JSON(customError.StatusCode, customError)
	}

	//Dto to product
	newProduct := productDto.ToProduct()
	//Insert to db
	if err := ds.GenericRepository.GenericProduct.Insert(newProduct); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(201, newProduct.Id)

}

func (ds *OrderService) PutProduct(c echo.Context) error {
	var productDto dto.ProductForUpdateDto
	//Binding
	if err := c.Bind(&productDto); err != nil {
		return c.JSON(customerror.InvalidBodyError.StatusCode, customerror.InvalidBodyError)
	}

	//Validation
	if err := validator.New().Struct(productDto); err != nil {
		customError := customerror.NewError(err.Error(), 400)
		return c.JSON(customError.StatusCode, customError)
	}

	//Dto to product
	updatedProduct := productDto.ToProduct()
	//Update
	if err := ds.GenericRepository.GenericProduct.Update(updatedProduct, updatedProduct.Id); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}

func (ds *OrderService) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	if err := ds.GenericRepository.GenericProduct.Delete(id); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}
