package users

import (
	"../../cmd"
	"../../utility"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func UserRouterInit(router *mux.Router) {
	router.Handle("/users", cmd.Middleware(http.HandlerFunc(GetUser), cmd.AuthMiddleware)).Methods(cmd.GET)
	//refresh token doesnt have userid
	router.HandleFunc("/users/refresh", RefreshUserToken).Methods(cmd.GET)
}

// GetVideos take all videos list.
func GetUser(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("userID")

	if len(userID) == 0 {
		utility.GenerateError(w, errors.New("User ID must be required"), http.StatusNotFound, "User ID Not found.")
		return
	}

	var ctx = context.Background()
	app := cmd.FBInstance()

	database, error := app.Firestore(ctx)
	if error != nil {
		utility.GenerateError(w, error, http.StatusInternalServerError, "")
		return
	}

	token, error := cmd.GetUserData(userID)
	if error != nil {
		return
	}

	document, error := database.Collection(cmd.FIRESTORE_USERS).Doc(token.UserID).Get(ctx)
	if error != nil {
		utility.GenerateError(w, error, http.StatusInternalServerError, "Firebase have error.")
		return
	}

	var user cmd.Users

	if err := document.DataTo(&user); err != nil {
		utility.GenerateError(w, error, http.StatusNotFound, "User id not found for the database")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// GetVideos take all videos list.
func RefreshUserToken(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get(cmd.QUERY_USER_ID)
	apiKey := r.Header.Get(cmd.QUERY_API_KEY)

	if len(apiKey) == 0 {
		error := errors.New("Api key required.")
		utility.GenerateError(w, error, http.StatusNotAcceptable, "")
		return
	}

	error := cmd.RefreshUserToken(userID)
	if error != nil {
		utility.GenerateError(w, error, http.StatusNotAcceptable, "User id not found for the database")
		return
	}

	var success utility.BaseSuccess
	success.CreatedDate = time.Now().String()
	success.Success = true
	success.Data = "User token refresh completed."

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)

}
