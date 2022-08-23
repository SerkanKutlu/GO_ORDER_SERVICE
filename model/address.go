package model

type Address struct {
	AddressLine string `bson:"addressLine" json:"addressLine" validate:"required"`
	City        string `bson:"city" json:"city" validate:"required"`
	Country     string `bson:"country" json:"country" validate:"required"`
	CityCode    int    `bson:"cityCode" json:"cityCode" validate:"gt=0"`
}
