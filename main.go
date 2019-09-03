package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Book struct {
	Title  string `json:title`
	Author string `json:author`
}

func init() {

	opt := option.WithCredentialsFile("assets/you2win.json")
	config := &firebase.Config{ProjectID: "you2win-3b9d9", DatabaseURL: "https://you2win-3b9d9.firebaseio.com/"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	defaultDatabase, err := app.Database(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)

	}
	_book := Book{Title: "Veli", Author: "Bacik"}
	data, err := defaultDatabase.NewRef("test").Child("testingo").Push(context.Background(), _book)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	print(data.Path)

	// defaultClient, err := app.Auth(context.Background())
	// if err != nil {
	// 	log.Fatalf("error getting Auth client: %v\n", err)
	// }

	// user, err := defaultClient.GetUserByEmail(context.Background(), "velibacik@gmail.com")
	// if err != nil {
	// 	log.Fatalf("error getting Auth client: %v\n", err)

	// }

	// print(user.Email)

}

func refreshToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {

	// http.HandleFunc("/", greet)
	// http.ListenAndServe(":8080", nil)
}
