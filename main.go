package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Book struct {
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
}

var _app *firebase.App
var ctx context.Context

func init() {

	opt := option.WithCredentialsFile("assets/you2win.json")
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

	router.HandleFunc("/videos", getVideos).Methods("GET")
	http.ListenAndServe(":8000", router)
}
func postVideos(w http.ResponseWriter, r *http.Request) {
	// defaultDatabase, err := _app.Database(context.Background())
	// if err != nil {
	// 	log.Fatalf("error getting Auth client: %v\n", err)

	// }
	// _book := Book{Title: "Veli", Author: "Bacik"}
	// data, err := defaultDatabase.NewRef("test").Child("testingo").Push(context.Background(), _book)
	// if err != nil {
	// 	log.Fatalf("error getting Auth client: %v\n", err)
	// }
}
func getVideos(w http.ResponseWriter, r *http.Request) {
	if len(r.Header.Get("veli")) <= 0 {
		http.Error(w, "Header token must be write", http.StatusNotAcceptable)
		return
	}
	db, err := _app.Database(context.Background())
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
