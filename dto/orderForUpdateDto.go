package dto

import (
	"github.com/SerkanKutlu/orderService/model"
	"time"
)

type OrderForUpdateDto struct {
	Id         string        `bson:"_id" json:"_id" validate:"required"`               //calculate auto
	Status     string        `bson:"status" json:"status" validate:"required"`         //from user
	Address    model.Address `bson:"address" json:"address" validate:"required"`       //get from customer
	ProductIds []string      `bson:"productIds" json:"productIds" validate:"required"` //from user

}

func (dto *OrderForUpdateDto) ToOrder() *model.Order {
	return &model.Order{
		Id:         dto.Id,
		CustomerId: "",
		Quantity:   0,
		Total:      0,
		Status:     dto.Status,
		Address:    dto.Address,
		ProductIds: dto.ProductIds,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Now().UTC(),
	}
}
