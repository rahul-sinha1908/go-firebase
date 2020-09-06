package ifirebase

import (
	"context"

	firebase "firebase.google.com/go"
)

//App Firebase App
var (
	App *firebase.App
)

//InitiateFirebase This call is used to initiate Firebase
//GOOGLE_APPLICATION_CREDENTIALS <- This env variable should contain json file
//Example -> export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/service-account-file.json"
func InitiateFirebase() error {
	var err error
	App, err = firebase.NewApp(context.Background(), nil)
	return err
}
