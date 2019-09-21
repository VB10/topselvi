package cmd

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// FBInstance get firebase.
func FBInstance() *firebase.App {
	opt := option.WithCredentialsFile("configs/you2win.json")
	config := &firebase.Config{ProjectID: "you2win-3b9d9", DatabaseURL: "https://you2win-3b9d9.firebaseio.com/"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}

//todo : remove it
func FirebaseInstance() (*firebase.App, context.Context,error) {
	opt := option.WithCredentialsFile("configs/you2win.json")
	background := context.Background()
	config := &firebase.Config{ProjectID: "you2win-3b9d9", DatabaseURL: "https://you2win-3b9d9.firebaseio.com/"}
	app, err := firebase.NewApp(background, config, opt)
	if err != nil {
		return  nil,nil,err
	}
	return app, background,err
}
