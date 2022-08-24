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
		return nil, customerror.InternalServerError
	}
	var orders []model.Order
	var order model.Order
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&order); err != nil {
			return nil, customerror.InternalServerError
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
		return customerror.InternalServerError
	}
	if result.ModifiedCount == 0 {
		return customerror.OrderNotFoundError
	}
	return nil
}

func (ms *MongoService) DeleteOrdersOfCustomer(customerId string) *customerror.CustomError {
	result, err := ms.Orders.DeleteMany(context.TODO(), bson.M{"customerId": customerId})
	if err != nil {
		return customerror.InternalServerError
	}
	if result.DeletedCount == 0 {
		return customerror.OrderNotFoundError
	}
	return nil
}
