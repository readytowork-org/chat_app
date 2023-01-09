package infrastructure

import (
	"context"
	"fmt"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

type Database struct {
	DB *db.Client
}

// NewFBApp creates new firebase app instance
func NewFBApp(logger Logger) *firebase.App {
	ctx := context.Background()
	serviceAccountKeyFilePath, err := filepath.Abs("./FirebaseSevicekey.json")
	if err != nil {
		logger.Zap.Panic("Unable to load serviceAccountKey.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	conf := &firebase.Config{
		DatabaseURL: "https://letschat-c3e14-default-rtdb.asia-southeast1.firebasedatabase.app",
	}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		logger.Zap.Fatalf("Firebase NewApp: %v", err)
	}
	logger.Zap.Info("✅ Firebase app initialized.")
	return app
}

// NewFBAuth creates new firebase auth client
func NewFBAuth(logger Logger, app *firebase.App) *auth.Client {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	firebaseAuth, err := app.Auth(ctx)
	if err != nil {
		logger.Zap.Fatal("Firebase Authentication: %v", err)
	}
	return firebaseAuth
}

func NewFBDatabase(logger Logger, app *firebase.App) Database {
	fmt.Print(app, "adaldflkajdflkajl")
	logger.Zap.Info(app)
	client, err := app.Database(context.Background())
	if err != nil {
		logger.Zap.Error("Error in creating FB database client", err)
	}
	return Database{
		DB: client,
	}
}
