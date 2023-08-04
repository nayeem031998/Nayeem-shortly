package database

import (
	"context"
	"urlshortnerService/constant"
	"urlshortnerService/types"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	result, err := inst.InsertOne(context.TODO(), data)
	return result.InsertedID, err
}

func (mgr *manager) GetUrlFromCode(code string, collectionName string) (resp types.UrlDb, err error) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	err = inst.FindOne(context.TODO(), bson.M{"url_code": code}).Decode(&resp)
	return resp, err
}
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("myDatabase").Collection("url")
	return collection
}
func (mgr *manager) UpdateClicks(code string, collectionName string) (resp types.UrlDb) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	inst.FindOneAndUpdate(context.TODO(), bson.M{"url_code": code}, bson.M{
		"$inc": bson.M{"clicks": 1},
	})
	return resp
}

func (mgr *manager) UpdateLocation (code string, collectionName string, location string) (resp types.UrlDb) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	inst.FindOneAndUpdate(context.TODO(), bson.M{"url_code": code}, bson.M{
		"$push": bson.M{"locations": location},
	})
	return resp
}
