package users

import (
	"../../cmd"
	"../../utility"
	"context"
	"encoding/json"
	"github.com/VB10/topselvi/pkg/auth"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

func UserRouterInit(router *mux.Router) {
	router.Handle("/users", auth.Middleware(http.HandlerFunc(GetUser), auth.AuthMiddleware)).Methods(cmd.GET)
	//router.Handle("/videos", auth.Middleware(http.HandlerFunc(PostVideo), auth.AuthMiddleware)).Methods(cmd.POST)
}

func checkUserToken(val string) {

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

	client, error := app.Auth(ctx)
	if error != nil {
		utility.GenerateError(w, error, http.StatusNotFound, "")
		return
	}

	xa, error := client.GetUser(ctx, userID)
	print(xa.DisplayName)

	token, error := client.CustomToken(ctx, xa.UID)
	if error != nil {
		utility.GenerateError(w, error, http.StatusNotFound, "")
		return
	}

	verifyToken, error := client.VerifyIDTokenAndCheckRevoked(ctx, token)
	if error != nil {
		utility.GenerateError(w, error, http.StatusNotFound, "")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(verifyToken)

}
