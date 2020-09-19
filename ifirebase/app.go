package ifirebase

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

//App Firebase App
var (
	App        *firebase.App
	configPath string
)

//InitiateFirebase This call is used to initiate Firebase
//GOOGLE_APPLICATION_CREDENTIALS <- This env variable should contain json file
//Example -> export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/service-account-file.json"
func InitiateFirebase() error {
	if configPath != "" {
		return InitiateFirebaseWithConfig(configPath)
	}
	var err error
	App, err = firebase.NewApp(context.Background(), nil)
	return err
}

//InitiateFirebaseWithConfig Initiate firebase with config file
func InitiateFirebaseWithConfig(path string) error {
	configPath = path
	opt := option.WithCredentialsFile(path)
	var err error
	App, err = firebase.NewApp(context.Background(), nil, opt)
	return err
}
