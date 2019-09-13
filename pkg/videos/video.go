package videos

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"../../cmd"
)

// Video model.
type Video struct {
	VideoURL   string `json:"videoUrl"`
	VideoTitle string `json:"videoTitle"`
	User       User   `json:"user"`
}

// User model.
type User struct {
	Username string `json:"username"`
	ProfileImage  string `json:"profileImage"`
}

// GetVideos take all videos list.
func GetVideos(w http.ResponseWriter, r *http.Request) {

	var ctx = context.Background()
	app := cmd.FBInstance()
	db, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return
	}

	docsnap, err := db.Collection("videos").Documents(ctx).GetAll()
	if err != nil {
	}

	// kvp key
	var x = YoutubeConfig("V_zfNjN32f4")


	var v Video
	var videos []Video
	for _, doc := range docsnap {
		if err := doc.DataTo(&v); err != nil {
			println(err)
		}
		v.User = *x
		videos = append(videos, v)
	}



	// fmt.Println(v)

	json.NewEncoder(w).Encode(videos)

}
