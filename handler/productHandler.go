package handler

import (
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/dto"
	"github.com/SerkanKutlu/orderService/model"
	"github.com/go-playground/validator/v10"
)

func (ds *OrderService) GetAllProducts() (*[]model.Product, *customerror.CustomError) {
	products, err := ds.GenericRepository.GenericProduct.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ds *OrderService) GetByIdProduct(id string) (*model.Product, *customerror.CustomError) {
	product, err := ds.GenericRepository.GenericProduct.FindById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ds *OrderService) PostProduct(productDto *dto.ProductForCreationDto) (*string, *customerror.CustomError) {

	//Validation
	if err := validator.New().Struct(productDto); err != nil {
		customError := customerror.NewError(err.Error(), 400)
		return nil, customError
	}

	//Dto to product
	newProduct := productDto.ToProduct()
	//Insert to db
	if err := ds.GenericRepository.GenericProduct.Insert(newProduct); err != nil {
		return nil, err
	}
	return &newProduct.Id, nil

}

func (ds *OrderService) PutProduct(productDto *dto.ProductForUpdateDto) (*model.Product, *customerror.CustomError) {

	//Validation
	if err := validator.New().Struct(productDto); err != nil {
		customError := customerror.NewError(err.Error(), 400)
		return nil, customError
	}
	//Dto to product
	updatedProduct := productDto.ToProduct()
	//Update
	if err := ds.GenericRepository.GenericProduct.Update(updatedProduct, updatedProduct.Id); err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (ds *OrderService) DeleteProduct(id string) *customerror.CustomError {
	if err := ds.GenericRepository.GenericProduct.Delete(id); err != nil {
		return err
	}
	return nil
}
