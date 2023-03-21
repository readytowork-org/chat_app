package infrastructure

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database modal
type Database struct {
	DB *mongo.Database
}

// NewDatabase creates a new database instance
func NewDatabase(env Env, zapLogger Logger) Database {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://chatapp:chatapp123@cluster0.g13tyt9.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)

	zapLogger.Zap.Info("Database connection established")
	return Database{
		DB: client.Database("chatApp"),
	}
}
