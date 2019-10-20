package main

import (
	"github.com/VB10/topselvi/pkg/users"
	"github.com/VB10/topselvi/pkg/videos"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	_ = godotenv.Load()

	videos.VideoRouterInit(router)
	users.UserRouterInit(router)

	_ = http.ListenAndServe(":8000", router)
}
