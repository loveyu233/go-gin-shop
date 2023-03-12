package sys

import "go.mongodb.org/mongo-driver/mongo"

type MongoClient struct {
	Db *mongo.Database
}

func (m *MongoClient) Like() *mongo.Collection {
	return m.Db.Collection("like")
}

func (m *MongoClient) Collect() *mongo.Collection {
	return m.Db.Collection("collect")
}

func (m *MongoClient) Comment() *mongo.Collection {
	return m.Db.Collection("comment")
}
