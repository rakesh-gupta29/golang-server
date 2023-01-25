package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/rakesh-gupta29/golang-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var databaseName string

func GetCollection(name string) *mongo.Collection {
	return mongoClient.Database(databaseName).Collection(name)
}

func ConnectToMongo() error {
	app_config, config_err := config.LoadConfig()
	if config_err != nil {
		fmt.Println("Error Loading config")
	}
	if app_config.DB_URL == "" {
		return errors.New("you must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	database := app_config.DB_NAME
	if database == "" {
		return errors.New("you must set your 'DATABASE' environmental variable")
	} else {
		databaseName = database
	}

	var err error
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(app_config.DB_URL))
	if err != nil {
		panic(err)
	}
	return nil
}

func CloseMongoDBConnection() {
	err := mongoClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}
