package imessaging

import (
	"context"

	"firebase.google.com/go/messaging"
	"github.com/rahul-sinha1908/go-firebase/ifirebase"
)

func getMessagingContext() (*messaging.Client, error) {
	if ifirebase.App == nil {
		err := ifirebase.InitiateFirebase()
		if err != nil {
			return nil, err
		}
	}
	ctx := context.Background()
	client, err := ifirebase.App.Messaging(ctx)
	return client, err
}

//SendMessage Send Message
func SendMessage(notification messaging.Notification, data map[string]string, registrationTokens []string) (*messaging.BatchResponse, error) {
	client, err := getMessagingContext()
	if err != nil {
		return nil, err
	}

	message := &messaging.MulticastMessage{
		Data:         data,
		Notification: &notification,
		Tokens:       registrationTokens,
	}

	response, err := client.SendMulticast(context.Background(), message)

	return response, err
}
