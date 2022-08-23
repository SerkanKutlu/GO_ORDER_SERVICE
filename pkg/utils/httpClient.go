package utils

import (
	"edu_src_Go/customerror"
	"edu_src_Go/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetCustomerAddress(customerId string) (*model.Address, *customerror.CustomError) {
	url := "http://localhost:5000/customers/" + customerId
	resp, err := http.Get(url)
	if err != nil {
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
