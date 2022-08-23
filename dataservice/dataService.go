package dataservice

import (
	"edu_src_Go/customerror"
	"edu_src_Go/model"
)

type OrderDataInterface interface {
	FindAllOrders() (*[]model.Order, *customerror.CustomError)
	FindByIdOrder(id string) (*model.Order, *customerror.CustomError)
	FindOrdersOfCustomer(customerId string) (*[]model.Order, *customerror.CustomError)
	InsertOrder(order *model.Order) *customerror.CustomError
	UpdateOrder(order *model.Order) *customerror.CustomError
	UpdateStatusFieldOrder(id string, newStatus string) *customerror.CustomError
	DeleteOrder(id string) *customerror.CustomError
	DeleteOrdersOfCustomer(customerId string) *customerror.CustomError
}

type ProductDataInterface interface {
	FindAllProducts() (*[]model.Product, *customerror.CustomError)
	FindByIdProduct(id string) (*model.Product, *customerror.CustomError)
	InsertProduct(product *model.Product) *customerror.CustomError
	UpdateProduct(product *model.Product) *customerror.CustomError
	DeleteProduct(id string) *customerror.CustomError
}