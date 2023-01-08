package infrastructure

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

func NewFBCredentials(logger Logger) (context.Context, option.ClientOption) {
	ctx := context.Background()
	serviceAccountKeyFilePath, err := filepath.Abs("./FirebaseSevicekey.json")
	if err != nil {
		logger.Zap.Panic("Unable to load serviceAccountKey.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	return ctx, opt
}

// NewFBApp creates new firebase app instance
func NewFBApp(logger Logger, ctx  opt option.ClientOption) *firebase.App {
	app, err := firebase.NewApp(ctx, nil, opt)
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
		logger.Zap.Fatalf("Firebase Authentication: %v", err)
	}
	return firebaseAuth
}

func NewDatabase(logger Logger, app *firebase.App) *db.Client {
	client, err := app.Database(context.Background())
	if err != nil {
		logger.Zap.Error("Error in creating FB database client", err)
	}
	return client
}
