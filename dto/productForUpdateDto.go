package dto

import (
	"edu_src_Go/model"
)

type ProductForUpdateDto struct {
	Id       string  `bson:"_id" json:"_id" validate:"required"`
	ImageUrl string  `bson:"imageUrl" json:"imageUrl" validate:"required"`
	Name     string  `bson:"name" json:"name" validate:"required"`
	Price    float32 `bson:"price" json:"price" validate:"required"`
}

func (dto *ProductForUpdateDto) ToProduct() *model.Product {
	return &model.Product{
		Id:       dto.Id,
		ImageUrl: dto.ImageUrl,
		Name:     dto.Name,
		Price:    dto.Price,
	}
}
