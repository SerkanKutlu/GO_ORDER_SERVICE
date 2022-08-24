package mongodb

import (
	"context"
	"github.com/SerkanKutlu/orderService/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoService struct {
	Orders   *mongo.Collection
	Products *mongo.Collection
}

func GetMongoService(mongoConfig *config.MongoConfig) *MongoService {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConfig.ConnectionString))
	if err != nil {
		panic(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	database := client.Database(mongoConfig.Database)
	productCollection := database.Collection(mongoConfig.Collection["product"])
	orderCollection := database.Collection(mongoConfig.Collection["order"])
	service := MongoService{
		orderCollection, productCollection,
	}
	return &service
}
