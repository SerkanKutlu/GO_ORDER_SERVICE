package mongodb

import (
	"context"
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Entity interface {
	model.Order | model.Product
}

type GenericRepository[T Entity] struct {
	Collection *mongo.Collection
}

func (ec *GenericRepository[T]) FindAll() (*[]T, *customerror.CustomError) {
	cursor, err := ec.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, customerror.InternalServerError
	}
	var entities []T
	var entity T
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&entity); err != nil {
			return nil, customerror.InternalServerError
		}
		entities = append(entities, entity)
	}
	return &entities, nil
}
func (ec *GenericRepository[T]) FindById(id string) (*T, *customerror.CustomError) {
	filter := bson.M{"_id": id}
	var entity T
	err := ec.Collection.FindOne(context.Background(), filter).Decode(&entity)
	if err != nil {
		return nil, customerror.OrderNotFoundError
	}
	return &entity, nil
}
func (ec *GenericRepository[T]) Insert(entity *T) *customerror.CustomError {
	_, err := ec.Collection.InsertOne(context.Background(), entity)
	if err != nil {
		return customerror.InternalServerError
	}
	return nil
}
func (ec *GenericRepository[T]) Update(entity *T, entityId string) *customerror.CustomError {
	result, err := ec.Collection.ReplaceOne(context.Background(), bson.M{"_id": entityId}, entity)
	if err != nil {
		return customerror.InternalServerError
	}
	if result.MatchedCount == 0 {
		return customerror.OrderNotFoundError
	}
	return nil
}
func (ec *GenericRepository[T]) Delete(id string) *customerror.CustomError {
	result, err := ec.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return customerror.InternalServerError
	}
	if result.DeletedCount == 0 {
		return customerror.OrderNotFoundError
	}
	return nil
}
