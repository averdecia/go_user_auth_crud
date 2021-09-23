package mongodb

import (
	"context"
	"crud/drivers/database"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoTimeOut time to wait from MongoDB
const MongoTimeOut = 30 * time.Second

// MongoID default field to save in mongodb
type MongoID struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
}

// MongoIDInterface interface to generate mongo id
type MongoIDInterface interface {
	GenerateID()
}

// GenerateID attach default user id before insert
func (s *MongoID) GenerateID() {
	s.ID = primitive.NewObjectID()
}

// GetByID function
func GetByID(client *database.MongoDBClient, coll string, id string, entity interface{}) (interface{}, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return GetOneBy(client, coll, bson.M{"_id": objectID}, entity)
}

// GetOneBy function
func GetOneBy(client *database.MongoDBClient, coll string, query bson.M, entity interface{}) (interface{}, error) {
	var ctx context.Context
	element := reflect.New(reflect.TypeOf(entity)).Interface()

	var collection = client.GetConnection().Collection(coll)
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeOut)
	defer cancel()
	err := collection.FindOne(ctx, query).Decode(element)
	if err != nil {
		fmt.Printf("error %v", err)
		return nil, err
	}
	return element, nil
}

// GetAll function
func GetAll(client *database.MongoDBClient, coll string, query bson.M, entity interface{}) ([]interface{}, error) {
	if query == nil {
		query = bson.M{}
	}
	var elements []interface{}
	var ctx context.Context
	var collection = client.GetConnection().Collection(coll)
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeOut)
	defer cancel()
	rows, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Printf("error %v", err)
		return nil, err
	}
	defer rows.Close(ctx)
	for rows.Next(ctx) {

		element := reflect.New(reflect.TypeOf(entity)).Interface()
		err := rows.Decode(element)

		if err != nil {
			fmt.Printf("error %v", err)
			return nil, err
		}
		elements = append(elements, element)
	}

	return elements, nil
}

// DeleteByID function
func DeleteByID(client *database.MongoDBClient, coll string, id string) (bool, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return DeleteAll(client, coll, bson.M{"_id": objectID})
}

// DeleteAll function
func DeleteAll(client *database.MongoDBClient, coll string, query bson.M) (bool, error) {
	var ctx context.Context

	var collection = client.GetConnection().Collection(coll)
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeOut)
	defer cancel()
	_, err := collection.DeleteMany(ctx, query)
	if err != nil {
		fmt.Printf("error %v", err)
		return false, err
	}
	return true, nil
}

// Create function
func Create(client *database.MongoDBClient, coll string, entity MongoIDInterface) (interface{}, error) {
	var ctx context.Context
	entity.GenerateID()

	var collection = client.GetConnection().Collection(coll)
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeOut)
	defer cancel()
	_, err := collection.InsertOne(ctx, entity)
	if err != nil {
		fmt.Printf("error %v", err)
		return false, err
	}
	return entity, nil
}

// UpdateOne function
func UpdateOne(client *database.MongoDBClient, coll string, query bson.M, setElements bson.D, entity interface{}) (interface{}, error) {
	var ctx context.Context
	var collection = client.GetConnection().Collection(coll)
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeOut)
	defer cancel()

	findOptions := &options.FindOneAndUpdateOptions{}
	findOptions.SetUpsert(false)
	findOptions.SetReturnDocument(options.After)

	err := collection.FindOneAndUpdate(ctx, query, bson.D{
		{Key: "$set", Value: setElements},
	}, findOptions).Decode(&entity)

	if err != nil {
		fmt.Printf("error %v", err)
		return false, err
	}
	return entity, nil
}

// UpdateAll function
func UpdateAll(client *database.MongoDBClient, coll string, query bson.M, setElements bson.D, entities []interface{}, entity interface{}) (int64, error) {
	var ctx context.Context
	var collection = client.GetConnection().Collection(coll)
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeOut)
	defer cancel()
	result, err := collection.UpdateMany(ctx, query, bson.D{
		{Key: "$set", Value: setElements},
	})

	if err != nil {
		fmt.Printf("error %v", err)
		return 0, err
	}
	return result.ModifiedCount, nil
}
