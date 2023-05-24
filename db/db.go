package db

import (
	"context"
	"os"

	"github.com/vinitparekh17/go-mongo/utility"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database
var Collection *mongo.Collection
var Ctx context.Context

func Init() {
	clientOpt := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	Ctx = context.Background()
	client, err := mongo.Connect(Ctx, clientOpt)
	utility.ErrorHandler(err)
	err = client.Ping(context.Background(), nil)
	utility.ErrorHandler(err)
	Database = client.Database(os.Getenv("MONGO_DB"))
	Collection = Database.Collection(os.Getenv("MONGO_COLLECTION"))
}

func GetDb() *mongo.Database {
	return Database
}
