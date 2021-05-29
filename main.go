package main

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EchobotLogtrace struct {
	Hostname  string
	Message   string
	Timestamp time.Time
}

type MongoDBConfig struct {
	Uri        string
	Database   string
	Collection string
}

var (
	outputType    string = "stdout"
	message              = "Gitops Flux series!"
	sleepTime     string = "5s"
	trace         EchobotLogtrace
	mongodbConfig *MongoDBConfig
	client        *mongo.Client
)

func NewMongoDBConfig() (*MongoDBConfig, error) {
	var uri, database, collection string

	if len(os.Getenv("MONGODB_URI")) != 0 {
		uri = os.Getenv("MONGODB_URI")
		log.Println("MongoDB uri detected")
	} else {
		return nil, errors.New("Missing MongoDB uri")
	}

	if len(os.Getenv("MONGODB_DATABASE")) != 0 {
		database = os.Getenv("MONGODB_DATABASE")
	} else {
		database = "echobot"
	}
	log.Printf("MongoDB database: %s", database)

	if len(os.Getenv("MONGODB_COLLECTION")) != 0 {
		collection = os.Getenv("MONGODB_COLLECTION")
	} else {
		collection = "log"
	}
	log.Printf("MongoDB collection: %s", collection)

	return &MongoDBConfig{uri, database, collection}, nil
}

func ExecMongoDB(c *mongo.Collection, t EchobotLogtrace) {
	// insert trace in database
	insertResult, err := c.InsertOne(context.TODO(), t)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Traza insertada en base de datos (%s): %s - %s\n", insertResult.InsertedID, t.Hostname, t.Message)
}

func ExecStdout(t EchobotLogtrace) {
	log.Printf("hostname: %s - %s\n", t.Hostname, t.Message)
}

func setConfig() (string, string, EchobotLogtrace, error) {
	// output to use
	if len(os.Getenv("OUTPUT_TYPE")) != 0 {
		outputType = os.Getenv("OUTPUT_TYPE")
	}

	// get hostname from os
	hostname, err := os.Hostname()
	if err != nil {
		return "", "", EchobotLogtrace{}, errors.New("Error getting hostname.")
	}

	// message to treat
	var message string
	if len(os.Getenv("MESSAGE")) != 0 {
		message = os.Getenv("MESSAGE")
	}

	// sleep time
	if len(os.Getenv("SLEEP_TIME")) != 0 {
		sleepTime = os.Getenv("SLEEP_TIME")
	}

	trace := EchobotLogtrace{hostname, message, time.Now()}

	return outputType, sleepTime, trace, nil
}

func main() {
	// check common variables
	outputType, sleepTime, trace, err := setConfig()
	if err != nil {
		log.Panicln("Error getting hostname.")
	}

	log.Printf("Started echobot in mode: %s", outputType)

	switch outputType {
	case "mongodb":
		// config mongo connection
		mongodbConfig, err = NewMongoDBConfig()
		if err != nil {
			log.Panicln(err)
		}

		// connect to MongoDB
		client, err = mongo.NewClient(options.Client().ApplyURI(mongodbConfig.Uri))
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		defer client.Disconnect(ctx)
	}

	for {
		// select mode
		switch outputType {
		case "mongodb":
			// Get collection
			collection := client.Database(mongodbConfig.Database).Collection(mongodbConfig.Collection)

			// Insert trace
			ExecMongoDB(collection, trace)
		default:
			ExecStdout(trace)
		}

		// rest a little
		sleepTimeDuration, _ := time.ParseDuration(sleepTime)
		time.Sleep(sleepTimeDuration)
	}
}
