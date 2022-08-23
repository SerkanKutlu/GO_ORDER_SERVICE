package mongodb

import (
	"context"
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/SerkanKutlu/orderService/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (md *MongoDb) FindAllProducts() (*[]model.Product, *customerror.CustomError) {
	cursor, err := md.Products.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, customerror.InternalServerError
	}
	var products []model.Product
	var product model.Product
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&product); err != nil {
			return nil, customerror.InternalServerError
		}
		products = append(products, product)
	}
	return &products, nil
}

func (md *MongoDb) FindByIdProduct(id string) (*model.Product, *customerror.CustomError) {
	filter := bson.M{"_id": id}
	var product model.Product
	err := md.Products.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return nil, customerror.ProductNotFoundError
	}
	return &product, nil
}

func (md *MongoDb) InsertProduct(product *model.Product) *customerror.CustomError {
	_, err := md.Products.InsertOne(context.Background(), product)
	if err != nil {
		return customerror.InternalServerError
	}
	return nil
}

func (md *MongoDb) UpdateProduct(product *model.Product) *customerror.CustomError {
	result, err := md.Products.ReplaceOne(context.Background(), bson.M{"_id": product.Id}, product)
	if err != nil {
		return customerror.InternalServerError
	}
	if result.MatchedCount == 0 {
		return customerror.ProductNotFoundError
	}
	return nil
}

func (md *MongoDb) DeleteProduct(id string) *customerror.CustomError {
	result, err := md.Products.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return customerror.InternalServerError
	}
	if result.DeletedCount == 0 {
		return customerror.ProductNotFoundError
	}
	return nil
}
