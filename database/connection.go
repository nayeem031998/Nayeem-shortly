package database

import (
	"context"
	"fmt"
	"log"
	//"os"
	"time"
	"urlshortnerService/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type manager struct {
	connection *mongo.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

var Mgr Manager

type Manager interface {
	Insert(interface{}, string) (interface{}, error)
	GetUrlFromCode(string, string) (types.UrlDb, error)
	UpdateClicks(string, string) types.UrlDb
}

func ConnectDb() {
	Mongo_URL := "mongodb+srv://nayeemakhtar371:Kamakazi%40567@clusternayeem.xuaxvgp.mongodb.net/?retryWrites=true&w=majority"
	//client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri)))
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))

	//client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri)))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	Mgr = &manager{connection: client, ctx: ctx, cancel: cancel}
}
func ConnectToDB() *mongo.Client {
	Mongo_URL := "mongodb+srv://nayeemakhtar371:Kamakazi%40567@clusternayeem.xuaxvgp.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}
	//err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	panic(err)
	// }
	 Mgr = &manager{connection: client, ctx: ctx, cancel: cancel}


	fmt.Println("Connected to mongoDB")
	return client
}
