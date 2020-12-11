package db

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/carlriis/Limitless-Lottery/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Get("MONGO_URL")))
	if err != nil {
		log.Fatalf("Error creating Mongo client: %s", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to Mongo: %s", err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Could not ping Mongo: %s", err)
	}

	log.Info("Looks like connection to Mongo was successfull")

	defer client.Disconnect(ctx)
}
