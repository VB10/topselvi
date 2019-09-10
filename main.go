package main

import (
	"net/http"

	"github.com/VB10/topselvi/pkg/auth"
	"github.com/gorilla/mux"

	"../topselvi/pkg/videos"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/videos", auth.Middleware(http.HandlerFunc(videos.GetVideos), auth.AuthMiddleware)).Methods("GET")

	http.ListenAndServe(":8000", router)
}
