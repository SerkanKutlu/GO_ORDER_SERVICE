package model

import "time"

type Order struct {
	Id         string    `bson:"_id" json:"_id" validate:"isdefault"`              //calculate auto
	CustomerId string    `bson:"customerId" json:"customerId" validate:"required"` //from user
	Quantity   int       `bson:"quantity" json:"quantity" validate:"isdefault"`    //calculate auto
	Total      float32   `bson:"total" json:"total" validate:"isdefault"`          //calculate auto
	Status     string    `bson:"status" json:"status" validate:"required"`         //from user
	Address    Address   `bson:"address" json:"address" validate:"isdefault"`      //get from customer
	ProductIds []string  `bson:"productIds" json:"productIds" validate:"required"` //from user
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt" validate:"isdefault"`  //calculate auto
	UpdatedAt  time.Time `bson:"updatedAt" json:"updatedAt" validate:"isdefault"`  //calculate auto
}
