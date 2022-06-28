package routes

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Database Conncetion Instance

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017/caloriesdb"

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Mongodb")
	return client
}

var Client *mongo.Client = DBinstance()

//To Add Collection to Our Database
func OpenCollection(cllient *mongo.Client, collectioonName string) *mongo.Collection {
	var collection *mongo.Collection = Client.Database("caloriesDb").Collection(collectioonName)
	return collection
}
