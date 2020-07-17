package ifirebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func initiateFirebase() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := []string{"YOUR_REGISTRATION_TOKEN"}

	// See documentation on defining a message payload.
	message := &messaging.MulticastMessage{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Notification: &messaging.Notification{
			Title:    "",
			Body:     "",
			ImageURL: "",
		},
		Tokens: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.SendMulticast(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response)
}
