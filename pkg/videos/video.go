package videos

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/VB10/topselvi/pkg/auth"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"time"

	"../../cmd"
	"../../utility"
)

// Video model.
type Videos struct {
	VideoURL        string `json:"videoUrl"`
	VideoTitle      string `json:"videoTitle"`
	User            User   `json:"user"`
	Price           int    `json:"price"`
	NumberOfMembers int    `json:"numberOfMembers"`
}

// User model.
type User struct {
	Username     string `json:"userName"`
	ProfileImage string `json:"profileImage"`
}

type VideoPost struct {
	YoutubeID       string `json:"youtubeID" validate:"required"`
	Price           int    `json:"price" validate:"required,min=1"`
	NumberOfMembers int    `json:"numberOfMembers" validate:"required,min=1"`
}

func VideosRouterInit(router *mux.Router) {
	router.Handle("/videos", auth.Middleware(http.HandlerFunc(GetVideos), auth.AuthMiddleware)).Methods(cmd.GET)
	router.Handle("/videos", auth.Middleware(http.HandlerFunc(PostVideo), auth.AuthMiddleware)).Methods(cmd.POST)
}

// GetVideos take all videos list.
func GetVideos(w http.ResponseWriter, r *http.Request) {

	var ctx = context.Background()
	app := cmd.FBInstance()
	db, error := app.Firestore(ctx)
	if error != nil {
		utility.GenerateError(w, error, http.StatusInternalServerError, "Firebase have error.")
		return
	}

	documents, error := db.Collection("videos").Documents(ctx).GetAll()
	if error != nil {
		utility.GenerateError(w, error, http.StatusInternalServerError, "Firebase have error.")
		return
	}

	var videos []Videos
	for _, doc := range documents {
		var v Videos
		if err := doc.DataTo(&v); err != nil {
			println(err)
		}
		videos = append(videos, v)
	}
	json.NewEncoder(w).Encode(videos)
}

// Post Videos take all videos list.
func PostVideo(w http.ResponseWriter, r *http.Request) {
	v := validator.New()
	var videoPost VideoPost

	json.NewDecoder(r.Body).Decode(&videoPost)
	error := v.Struct(videoPost)
	if error != nil {
		utility.GenerateError(w, error, http.StatusUnprocessableEntity, cmd.MODEL_INVALID)
		return
	}

	youtubeVideo, error := YoutubeVideoDetail(videoPost.YoutubeID)
	if error != nil {
		utility.GenerateError(w, error, http.StatusNotFound, "")
		return
	}

	youtubeUser, error := YoutubeUserDetail(videoPost.YoutubeID)
	if error != nil {
		utility.GenerateError(w, error, http.StatusNotFound, "")
		return
	}

	var videos Videos
	videos.VideoTitle = youtubeVideo.Snippet.Title
	videos.NumberOfMembers = videoPost.NumberOfMembers
	videos.Price = videoPost.Price
	videos.VideoURL = cmd.YOUTUBE_WATCH_PREFIX + videoPost.YoutubeID
	videos.User = *youtubeUser
	firestoreRef, error := videos.writeFirebaseDatabase()

	if error != nil {
		utility.GenerateError(w, error, http.StatusInternalServerError, "Firebase server have problem.")
		return
	}

	var success utility.BaseSuccess
	success.CreatedDate = time.Now().String()
	success.Success = true
	success.Data = firestoreRef.ID

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)

}

func (videos Videos) writeFirebaseDatabase() (*firestore.DocumentRef, error) {
	var ctx = context.Background()
	app := cmd.FBInstance()
	firestore, error := app.Firestore(ctx)
	if error != nil {
		log.Fatalf("error getting Auth client: %v\n", error)
		return nil, error
	}

	response, _, error := firestore.Collection(cmd.FIRESTORE_VIDEOS).Add(ctx, videos)
	if error != nil {
		return nil, error
	}
	return response, nil
}
