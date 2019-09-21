package users

import (
	"../../cmd"
	"../../utility"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func UserRouterInit(router *mux.Router) {
	router.Handle("/users", cmd.Middleware(http.HandlerFunc(GetUser), cmd.AuthMiddleware)).Methods(cmd.GET)
	router.HandleFunc("/users/refresh", RefreshUserToken).Methods(cmd.GET)
}

// GetVideos take all videos list.
func GetUser(w http.ResponseWriter, r *http.Request) {

	userToken := r.Header.Get(cmd.QueryUserToken)

	if len(userToken) == 0 {
		utility.GenerateError(w, errors.New("User ID must be required"), http.StatusNotFound, "User ID Not found.")
		return
	}

	var ctx = context.Background()
	app := cmd.FBInstance()

	database, err := app.Firestore(ctx)
	if err != nil {
		utility.GenerateError(w, err, http.StatusInternalServerError, "")
		return
	}

	token, err := cmd.GetUserData(userToken)
	if err != nil {
		return
	}

	document, err := database.Collection(cmd.FirestoreUsers).Doc(token.UserID).Get(ctx)
	if err != nil {
		utility.GenerateError(w, err, http.StatusInternalServerError, "Firebase have error.")
		return
	}

	var user cmd.Users
	if err := document.DataTo(&user); err != nil {
		utility.GenerateError(w, err, http.StatusNotFound, "User id not found for the database")
		return
	}
	claims, err := cmd.JWTParser(userToken)
	userID := fmt.Sprintf("%v", claims[cmd.FbUid])
	user.UserID = userID

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}

// GetVideos take all videos list.
func RefreshUserToken(w http.ResponseWriter, r *http.Request) {
	userToken := r.Header.Get(cmd.QueryUserToken)
	apiKey := r.Header.Get(cmd.QueryApiKey)

	if len(apiKey) == 0 {
		err := errors.New("Api key required.")
		utility.GenerateError(w, err, http.StatusNotAcceptable, "")
		return
	}

	newToken, err := cmd.RefreshUserToken(userToken)
	if err != nil {
		utility.GenerateError(w, err, http.StatusNotAcceptable, "User id not found for the database")
		return
	}

	var success utility.BaseSuccess
	success.CreatedDate = time.Now().String()
	success.Success = true
	success.Data = newToken

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(success)
}
