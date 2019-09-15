package main

import (
	"../topselvi/pkg/videos"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	godotenv.Load()
	videos.VideosRouterInit(router)

	http.ListenAndServe(":8000", router)
}
