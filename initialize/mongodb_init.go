package initialize

import (
	"context"
	"github.com/spf13/viper"
	"go-gin-shop/enter/sys"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *sys.MongoClient

func InitMongoDb(url string) *sys.MongoClient {
	var err error

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		return nil
	}

	mongoClient = &sys.MongoClient{
		Db: client.Database(viper.GetString("mongo.datasource")),
	}
	return mongoClient
}
