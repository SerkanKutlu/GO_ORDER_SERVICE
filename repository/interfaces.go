package repository

//
//import (
//	"github.com/SerkanKutlu/orderService/customerror"
//	"github.com/SerkanKutlu/orderService/model"
//)
//
//type IOrderRepository interface {
//	FindAllOrders() (*[]model.Order, *customerror.CustomError)
//	FindByIdOrder(id string) (*model.Order, *customerror.CustomError)
//	FindOrdersOfCustomer(customerId string) (*[]model.Order, *customerror.CustomError)
//	InsertOrder(order *model.Order) *customerror.CustomError
//	UpdateOrder(order *model.Order) *customerror.CustomError
//	UpdateStatusFieldOrder(id string, newStatus string) *customerror.CustomError
//	DeleteOrder(id string) *customerror.CustomError
//	DeleteOrdersOfCustomer(customerId string) *customerror.CustomError
//}
//
//type IProductRepository interface {
//	FindAllProducts() (*[]model.Product, *customerror.CustomError)
//	FindByIdProduct(id string) (*model.Product, *customerror.CustomError)
//	InsertProduct(product *model.Product) *customerror.CustomError
//	UpdateProduct(product *model.Product) *customerror.CustomError
//	DeleteProduct(id string) *customerror.CustomError
//}
