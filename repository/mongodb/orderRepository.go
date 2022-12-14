package mongodb

import (
	"context"
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (ms *MongoService) FindOrdersOfCustomer(customerId string) (*[]model.Order, *customerror.CustomError) {
	filter := bson.M{"customerId": customerId}
	cursor, err := ms.Orders.Find(context.Background(), filter)
	if err != nil {
		ce := customerror.NewError(err.Error(), 500)
		return nil, ce
	}
	var orders []model.Order
	var order model.Order
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&order); err != nil {
			ce := customerror.NewError(err.Error(), 500)
			return nil, ce
		}
		orders = append(orders, order)
	}
	return &orders, nil

}

func (ms *MongoService) UpdateStatusFieldOrder(id string, newStatus string) *customerror.CustomError {
	filter := bson.M{"_id": id}
	updateCommand := bson.D{
		{"$set", bson.D{{"status", newStatus}, {"updatedAt", time.Now().UTC()}}},
	}
	result, err := ms.Orders.UpdateOne(context.Background(), filter, updateCommand)
	if err != nil {
		ce := customerror.NewError(err.Error(), 500)
		return ce
	}
	if result.ModifiedCount == 0 {
		return customerror.NotFoundError
	}
	return nil
}

func (ms *MongoService) DeleteOrdersOfCustomer(customerId string) *customerror.CustomError {
	result, err := ms.Orders.DeleteMany(context.TODO(), bson.M{"customerId": customerId})
	if err != nil {
		ce := customerror.NewError(err.Error(), 500)
		return ce
	}
	if result.DeletedCount == 0 {
		return customerror.NotFoundError
	}
	return nil
}
