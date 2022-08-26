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

type GenericCollection[T Entity] struct {
	Collection *mongo.Collection
}

type GenericRepository struct {
	GenericOrder   *GenericCollection[model.Order]
	GenericProduct *GenericCollection[model.Product]
}

func (gc *GenericCollection[T]) FindAll() (*[]T, *customerror.CustomError) {
	cursor, err := gc.Collection.Find(context.Background(), bson.M{})
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
func (gc *GenericCollection[T]) FindById(id string) (*T, *customerror.CustomError) {
	filter := bson.M{"_id": id}
	var entity T
	err := gc.Collection.FindOne(context.Background(), filter).Decode(&entity)
	if err != nil {
		return nil, customerror.NotFoundError
	}
	return &entity, nil
}
func (gc *GenericCollection[T]) Insert(entity *T) *customerror.CustomError {
	_, err := gc.Collection.InsertOne(context.Background(), entity)
	if err != nil {
		return customerror.InternalServerError
	}
	return nil
}
func (gc *GenericCollection[T]) Update(entity *T, entityId string) *customerror.CustomError {
	result, err := gc.Collection.ReplaceOne(context.Background(), bson.M{"_id": entityId}, entity)
	if err != nil {
		return customerror.InternalServerError
	}
	if result.MatchedCount == 0 {
		return customerror.NotFoundError
	}
	return nil
}
func (gc *GenericCollection[T]) Delete(id string) *customerror.CustomError {
	result, err := gc.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return customerror.InternalServerError
	}
	if result.DeletedCount == 0 {
		return customerror.NotFoundError
	}
	return nil
}
