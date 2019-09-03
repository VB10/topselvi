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

func init() {
	// firebase init
	opt := option.WithCredentialsFile("path/to/refreshToken.json")
	config := &firebase.Config{ProjectID: "my-project-id"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
}

func refreshToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
