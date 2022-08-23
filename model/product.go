package model

type Product struct {
	Id       string  `bson:"_id" json:"_id" validate:"isdefault"`
	ImageUrl string  `bson:"imageUrl" json:"imageUrl" validate:"required"`
	Name     string  `bson:"name" json:"name" validate:"required"`
	Price    float32 `bson:"price" json:"price" validate:"required"`
}
