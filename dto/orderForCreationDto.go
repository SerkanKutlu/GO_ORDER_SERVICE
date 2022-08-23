package dto

import (
	"edu_src_Go/model"
	uuid "github.com/satori/go.uuid"
	"time"
)

type OrderForCreationDto struct {
	CustomerId string   `bson:"customerId" json:"customerId" validate:"required"` //calculate auto
	Status     string   `bson:"status" json:"status" validate:"required"`         //from user
	ProductIds []string `bson:"productIds" json:"productIds" validate:"required"` //from user

}

func (dto *OrderForCreationDto) ToOrder() *model.Order {
	return &model.Order{
		Id:         uuid.NewV4().String(),
		CustomerId: dto.CustomerId,
		Quantity:   0,
		Total:      0,
		Status:     dto.Status,
		Address:    model.Address{},
		ProductIds: dto.ProductIds,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Time{},
	}
}
