package db

import (
	"context"
	"time"

	"github.com/carlriis/Limitless-Lottery/config"
	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var usersCollection *mongo.Collection
var sessionIdentitiesCollection *mongo.Collection

func Connect() {

	createClient()

	ctx, cancel := getContext()
	defer cancel()

	err := client.Connect(ctx)

	if err != nil {
		log.Fatalf("Error connecting to Mongo: %s", err)
	}

	err = Ping(ctx)
	if err != nil {
		log.Fatalf("Could not ping Mongo: %s", err)
	}

	createUsersCollection()
	createSessionIdentities()

	log.Info("Connection to Mongo was successfull")
}

func Disconnect() {
	ctx, cancel := getContext()
	defer cancel()
	client.Disconnect(ctx)
}

func Ping(ctx context.Context) error {
	err := client.Ping(ctx, readpref.Primary())
	return err
}

func createUsersCollection() {
	usersCollection = client.Database("limitless_lottery").Collection("users")
}

func createSessionIdentities() {
	sessionIdentitiesCollection = client.Database("limitless_lottery").Collection("session_identities")
}

func createClient() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(config.Get("MONGO_URL")))
	if err != nil {
		log.Fatalf("Error creating Mongo client: %s", err)
	}
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
