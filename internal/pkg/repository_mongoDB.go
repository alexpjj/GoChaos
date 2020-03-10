package pkg

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
MongoRepository is an abstraction of utilities for CRUD operation on MongoDB
*/
type MongoRepository struct {
	client   *mongo.Client
	database string
}

/*
NewMongoRepository initializes the type mongoRepository
*/
func NewMongoRepository(host string, port string, database string) (MongoRepository, error) {
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		return MongoRepository{}, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return MongoRepository{}, err
	}

	return MongoRepository{client, database}, nil
}

func (repo *MongoRepository) getCollection(coll string) *mongo.Collection {
	return repo.client.Database(repo.database).Collection(coll)
}

/*
Get gets a single document
*/
func (repo *MongoRepository) Get(collection string, key string, value string) (interface{}, error) {
	var result interface{}
	filter := bson.D{primitive.E{Key: key, Value: value}}

	err := repo.getCollection(collection).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
Update updates a single document
*/
func (repo *MongoRepository) Update(collection string, key string, value string, doc interface{}) error {
	filter := bson.D{primitive.E{Key: key, Value: value}}

	_, err := repo.getCollection(collection).UpdateOne(context.TODO(), filter, bson.D{
		primitive.E{Key: "$set", Value: doc}})

	if err != nil {
		return err
	}
	return nil
}

/*
Create creates a single document
*/
func (repo *MongoRepository) Create(collection string, doc interface{}) error {
	_, err := repo.getCollection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}
	return nil
}

/*
Remove removes a single document
*/
func (repo *MongoRepository) Remove(collection string, key string, value string) error {
	filter := bson.D{primitive.E{Key: key, Value: value}}

	_, err := repo.getCollection(collection).DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}
	return nil
}
