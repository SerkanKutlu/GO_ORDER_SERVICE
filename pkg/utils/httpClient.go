package utils

import (
	"encoding/json"
	"fmt"
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/model"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	ServiceUrlMap map[string]string
}

func (hc *HttpClient) GetCustomerAddress(customerId string) (*model.Address, *customerror.CustomError) {
	url := fmt.Sprintf("%s%s%s", hc.ServiceUrlMap["CustomerService"], "customers/", customerId)
	resp, err := http.Get(url)
	if resp == nil {
		return nil, customerror.InternalServerError
	}
	if resp.StatusCode == 404 {
		return nil, customerror.CustomerNotFoundError
	}

	var address *model.Order
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, customerror.InternalServerError
	}
	if err := json.Unmarshal(body, &address); err != nil {
		return nil, customerror.AddressNotFoundError
	}
	return &address.Address, nil

}
