package handler

import (
	"edu_src_Go/customerror"
	"edu_src_Go/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (ds *DataAccessService) GetAllProducts(c echo.Context) error {
	products, err := ds.ProductService.FindAllProducts()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, products)
}

func (ds *DataAccessService) GetByIdProduct(c echo.Context) error {
	id := c.Param("id")
	product, err := ds.ProductService.FindByIdProduct(id)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, product)
}

func (ds *DataAccessService) PostProduct(c echo.Context) error {
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
	if err := ds.ProductService.InsertProduct(newProduct); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(201, newProduct.Id)

}

func (ds *DataAccessService) PutProduct(c echo.Context) error {
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
	if err := ds.ProductService.UpdateProduct(updatedProduct); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}

func (ds *DataAccessService) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	if err := ds.ProductService.DeleteProduct(id); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, "")
}
