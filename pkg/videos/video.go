package videos

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"time"

	"../../cmd"
	"../../utility"
	"../users"
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

func VideoRouterInit(router *mux.Router) {
	router.Handle("/videos", cmd.Middleware(http.HandlerFunc(GetVideos), cmd.AuthMiddleware)).Methods(cmd.GET)
	router.Handle("/videos", cmd.Middleware(http.HandlerFunc(PostVideo), cmd.AuthMiddleware)).Methods(cmd.POST)
}

// GetVideos take all videos list.
func GetVideos(w http.ResponseWriter, r *http.Request) {

	var ctx = context.Background()
	app := cmd.FBInstance()
	db, err := app.Firestore(ctx)
	if err != nil {
		utility.GenerateError(w, err, http.StatusInternalServerError, "Firebase have error.")
		return
	}

	documents, err := db.Collection("videos").Documents(ctx).GetAll()
	if err != nil {
		utility.GenerateError(w, err, http.StatusInternalServerError, "Firebase have error.")
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
	_ = json.NewEncoder(w).Encode(videos)
}

// Post Videos take all videos list.
func PostVideo(w http.ResponseWriter, r *http.Request) {
	v := validator.New()
	var videoPost VideoPost

	_ = json.NewDecoder(r.Body).Decode(&videoPost)
	err := v.Struct(videoPost)
	if err != nil {
		utility.GenerateError(w, err, http.StatusUnprocessableEntity, cmd.ModelInvalid)
		return
	}

	youtubeVideo, err := YoutubeVideoDetail(videoPost.YoutubeID)
	if err != nil {
		utility.GenerateError(w, err, http.StatusNotFound, "")
		return
	}

	youtubeUser, err := YoutubeUserDetail(videoPost.YoutubeID)
	if err != nil {
		utility.GenerateError(w, err, http.StatusNotFound, "")
		return
	}

	var videos Videos
	videos.VideoTitle = youtubeVideo.Snippet.Title
	videos.NumberOfMembers = videoPost.NumberOfMembers
	videos.Price = videoPost.Price
	videos.VideoURL = cmd.YoutubeWatchPrefix + videoPost.YoutubeID
	videos.User = *youtubeUser
	firestoreRef, err := videos.writeFirebaseDatabase(r.Header.Get(cmd.QueryUserId))

	if err != nil {
		utility.GenerateError(w, err, http.StatusInternalServerError, "Firebase server have problem.")
		return
	}

	var success utility.BaseSuccess
	success.CreatedDate = time.Now().String()
	success.Success = true
	success.Data = firestoreRef.ID

	users.GetUser(w, r)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(success)

}

func (videos Videos) writeFirebaseDatabase(uid string) (*firestore.DocumentRef, error) {
	var ctx = context.Background()
	app := cmd.FBInstance()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return nil, err
	}

	user, err := cmd.GetUserData(uid)
	if user.Wallet < videos.Price*videos.NumberOfMembers {
		//TODO: Multi Language
		return nil, errors.New("You don't have enough money.")
	}

	response, _, err := client.Collection(cmd.FirestoreVideos).Add(ctx, videos)
	if err != nil {
		return nil, err
	}
	return response, nil
}
