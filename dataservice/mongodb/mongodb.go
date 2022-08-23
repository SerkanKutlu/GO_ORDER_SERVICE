package mongodb

import (
	"context"
	"edu_src_Go/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoDb struct {
	Orders   *mongo.Collection
	Products *mongo.Collection
}

func GetMongoDb(mongoConfig *config.MongoConfig) *MongoDb {
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
	service := MongoDb{
		orderCollection, productCollection,
	}
	return &service
}