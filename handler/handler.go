package handler

import "edu_src_Go/dataservice"

type DataAccessService struct {
	OrderService   dataservice.OrderDataInterface
	ProductService dataservice.ProductDataInterface
}

var dataAccessService = new(DataAccessService)

func GetDataServices() *DataAccessService {
	return dataAccessService
}

func SetOrderDataService(orderService dataservice.OrderDataInterface) {
	dataAccessService.OrderService = orderService
}

func SetProductDataService(productService dataservice.ProductDataInterface) {
	dataAccessService.ProductService = productService
}
