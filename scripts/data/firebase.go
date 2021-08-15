package main

import (
	"context"
	"fmt"
	"os"

	"firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseService struct {
	app *firebase.App
}

func NewFirebaseService(clientSecretPath string) (*FirebaseService, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	app, err := firebase.NewApp(
		context.Background(),
		&firebase.Config{ProjectID: "mentu-lxs"},
		option.WithCredentialsFile(fmt.Sprintf("%s%s", path, clientSecretPath)),
	)
	if err != nil {
		return nil, err
	}

	return &FirebaseService{app: app}, err
}
