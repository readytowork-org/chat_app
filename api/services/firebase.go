package services

import (
	"context"
	"letschat/infrastructure"
	"log"

	"firebase.google.com/go/messaging"
)

type FirebaseService struct {
	fb     infrastructure.FirebaseApp
	logger infrastructure.Logger
}

func NewFirebaseService(
	fb infrastructure.FirebaseApp,
	logger infrastructure.Logger,
) FirebaseService {
	return FirebaseService{
		fb:     fb,
		logger: logger,
	}
}

func (fs FirebaseService) PushNotification(body string, title string, deviceId string) error {

	client, err := fs.fb.FA.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error initializing Messaging client: %v\n", err)
		return err
	}
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: deviceId,
	}
	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalf("error sending message: %v\n", err)
		return err
	}
	log.Printf("message sent successfully: %v\n", response)
	return nil
}
