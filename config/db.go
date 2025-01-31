package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	InfoLogger *log.Logger
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

func HandlePoolMonitor(evt *event.PoolEvent) {
	switch evt.Type {
	case event.PoolClosedEvent, event.PoolCleared:
		log.Println(evt)
	}

}

func ConnectDB() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	//client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	monitor := &event.PoolMonitor{
		Event: HandlePoolMonitor,
	}

	client, err := mongo.NewClient(
		options.Client().
			ApplyURI(os.Getenv("MONGO_URI")).
			SetConnectTimeout(3 * time.Second).
			SetPoolMonitor(monitor))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("DB")),
	}
}
