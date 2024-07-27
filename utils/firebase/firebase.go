package utils

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	ctx          context.Context
	FirebaseAuth *firebase.Auth
)

const (
	firebaseConfigFile = "./../../service-account-key.json"
)

func FirebaseInit() {

	ctx = context.Background()
	sa := option.WithCredentialsFile(firebaseConfigFile)

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	FirebaseAuth = client

}
