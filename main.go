package main

import (
	"./pkg/users"
	"./pkg/videos"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	godotenv.Load()
	videos.VideosRouterInit(router)
	users.UserRouterInit(router)

	http.ListenAndServe(":8000", router)
}
