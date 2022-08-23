package mongodb

import (
	"context"
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (md *MongoDb) FindAllOrders() (*[]model.Order, *customerror.CustomError) {
	cursor, err := md.Orders.Find(context.Background(), bson.M{})
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

func (md *MongoDb) FindByIdOrder(id string) (*model.Order, *customerror.CustomError) {
	filter := bson.M{"_id": id}
	var order model.Order
	err := md.Orders.FindOne(context.Background(), filter).Decode(&order)
	if err != nil {
		return nil, customerror.OrderNotFoundError
	}
	return &order, nil
}

func (md *MongoDb) FindOrdersOfCustomer(customerId string) (*[]model.Order, *customerror.CustomError) {
	filter := bson.M{"customerId": customerId}
	cursor, err := md.Orders.Find(context.Background(), filter)
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

func (md *MongoDb) InsertOrder(order *model.Order) *customerror.CustomError {
	_, err := md.Orders.InsertOne(context.Background(), order)
	if err != nil {
		return customerror.InternalServerError
	}
	return nil
}

func (md *MongoDb) UpdateOrder(order *model.Order) *customerror.CustomError {
	result, err := md.Orders.ReplaceOne(context.Background(), bson.M{"_id": order.Id}, order)
	if err != nil {
		return customerror.InternalServerError
	}
	if result.MatchedCount == 0 {
		return customerror.OrderNotFoundError
	}
	return nil
}

func (md *MongoDb) UpdateStatusFieldOrder(id string, newStatus string) *customerror.CustomError {
	filter := bson.M{"_id": id}
	updateCommand := bson.D{
		{"$set", bson.D{{"status", newStatus}, {"updatedAt", time.Now().UTC()}}},
	}
	result, err := md.Orders.UpdateOne(context.Background(), filter, updateCommand)
	if err != nil {
		return customerror.InternalServerError
	}
	if result.ModifiedCount == 0 {
		return customerror.OrderNotFoundError
	}
	return nil
}

func (md *MongoDb) DeleteOrder(id string) *customerror.CustomError {
	result, err := md.Orders.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return customerror.InternalServerError
	}
	if result.DeletedCount == 0 {
		return customerror.OrderNotFoundError
	}
	return nil
}

func (md *MongoDb) DeleteOrdersOfCustomer(customerId string) *customerror.CustomError {
	result, err := md.Orders.DeleteMany(context.TODO(), bson.M{"customerId": customerId})
	if err != nil {
		return customerror.InternalServerError
	}
	if result.DeletedCount == 0 {
		return customerror.OrderNotFoundError
	}
	return nil
}
