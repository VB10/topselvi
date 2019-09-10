package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"firebase.google.com/go/auth"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

// Book mock data
type Book struct {
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
}

var _app *firebase.App
var ctx context.Context

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}

func init() {

	opt := option.WithCredentialsFile("configs/you2win.json")
	config := &firebase.Config{ProjectID: "you2win-3b9d9", DatabaseURL: "https://you2win-3b9d9.firebaseio.com/"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	_app = app
	ctx = context.Background()
}

func refreshToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	router := mux.NewRouter()

	router.Handle("/videos", Middleware(http.HandlerFunc(exampleHandler), auth.AuthMiddleware)).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func getVideos(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("usertoken") != "" {
		http.Error(w, json.NewEncoder(w).Encode("a").Error(), http.StatusNotAcceptable)
		return
	}

	// cmd.verifyUser(_app)

	// client, err := _app.Auth(ctx)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	// token, err := client.GetUser(ctx, r.Header.Get("veli"))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	db, err := _app.Database(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return
	}

	ref := db.NewRef("test/testingo")
	results, err := ref.OrderByKey().GetOrdered(ctx)

	books := []Book{}
	if err != nil {
		log.Fatalln("Error querying database:", err)
	}
	for _, r := range results {
		// data := r.Unmarshal(Book)
		// books = append(books)
		var book Book

		if err := r.Unmarshal(&book); err != nil {
			print(err)
		}
		books = append(books, book)

	}
	json.NewEncoder(w).Encode(books)
}
