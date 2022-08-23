package dto

import (
	"edu_src_Go/model"
	uuid "github.com/satori/go.uuid"
)

type ProductForCreationDto struct {
	ImageUrl string  `bson:"imageUrl" json:"imageUrl" validate:"required"`
	Name     string  `bson:"name" json:"name" validate:"required"`
	Price    float32 `bson:"price" json:"price" validate:"required"`
}

func (dto *ProductForCreationDto) ToProduct() *model.Product {
	return &model.Product{
		Id:       uuid.NewV4().String(),
		ImageUrl: dto.ImageUrl,
		Name:     dto.Name,
		Price:    dto.Price,
	}
}
